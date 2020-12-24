#Analyse API server start process

## find func which exec start cmd 

~~~go
//cmd/dockerd/docker.go:71
func main() {
...
	cmd, err := newDaemonCommand()
...
}

//cmd/dockerd/docker.go:23
func newDaemonCommand() (*cobra.Command, error) {
...
	cmd := &cobra.Command{
		Use:           "dockerd [OPTIONS]",
		Short:         "A self-sufficient runtime for containers.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.flags = cmd.Flags()
			return runDaemon(opts)
		},
		DisableFlagsInUseLine: true,
		Version:               fmt.Sprintf("%s, build %s", dockerversion.Version, dockerversion.GitCommit),
	}
	cli.SetupRootCommand(cmd)
...
}

//cmd/dockerd/docker_unix.go:11
func runDaemon(opts *daemonOptions) error {
	daemonCli := NewDaemonCli()
	return daemonCli.start(opts)
}

//cmd/dockerd/daemon.go:77
func (cli *DaemonCli) start(opts *daemonOptions) (err error) {
...
	serverConfig, err := newAPIServerConfig(cli)
	if err != nil {
		return errors.Wrap(err, "failed to create API server")
	}
	cli.api = apiserver.New(serverConfig)
...
	serveAPIWait := make(chan error)
	go cli.api.Wait(serveAPIWait)
...
}

//api/server/server.go:198
func (s *Server) Wait(waitChan chan error) {
	if err := s.serveAPI(); err != nil {
		logrus.Errorf("ServeAPI error: %v", err)
		waitChan <- err
		return
	}
	waitChan <- nil
}
~~~

## 分析apiserver具体内容

* apiserver实例化: cli.api = apiserver.New(serverConfig)
* apiserver与http handler绑定: 
~~~go
//api/server/server.go:198
func (s *Server) Wait(waitChan chan error) {
	if err := s.serveAPI(); err != nil {
		logrus.Errorf("ServeAPI error: %v", err)
		waitChan <- err
		return
	}
	waitChan <- nil
}

//api/server/server.go:79
func (s *Server) serveAPI() error {
...
		srv.srv.Handler = s.createMux()
...
}

//api/server/server.go:166
func (s *Server) createMux() *mux.Router {
...
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			f := s.makeHTTPHandler(r.Handler())

			logrus.Debugf("Registering %s, %s", r.Method(), r.Path())
			m.Path(versionMatcher + r.Path()).Methods(r.Method()).Handler(f)
			m.Path(r.Path()).Methods(r.Method()).Handler(f)
		}
	}
...
}
~~~

s.routers对应的是路由，即功能列表，追查其初始化过程

~~~go
//api/server/server.go:153
func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
}
~~~

s.routers分两部分，一部分是s.routers，另一部分是routers...

##追查routers...
~~~go
//cmd/dockerd/daemon.go:470
func initRouter(opts routerOptions) {
...
	routers := []router.Router{
		// we need to add the checkpoint router before the container router or the DELETE gets masked
		checkpointrouter.NewRouter(opts.daemon, decoder),
		container.NewRouter(opts.daemon, decoder, opts.daemon.RawSysInfo(true).CgroupUnified),
		image.NewRouter(opts.daemon.ImageService()),
		systemrouter.NewRouter(opts.daemon, opts.cluster, opts.buildkit, opts.features),
		volume.NewRouter(opts.daemon.VolumesService()),
		build.NewRouter(opts.buildBackend, opts.daemon, opts.features),
		sessionrouter.NewRouter(opts.sessionManager),
		swarmrouter.NewRouter(opts.cluster),
		pluginrouter.NewRouter(opts.daemon.PluginManager()),
		distributionrouter.NewRouter(opts.daemon.ImageService()),
	}
...
	if len(grpcBackends) > 0 {
		routers = append(routers, grpcrouter.NewRouter(grpcBackends...))
	}

	if opts.daemon.NetworkControllerEnabled() {
		routers = append(routers, network.NewRouter(opts.daemon, opts.cluster))
	}
...
	opts.api.InitRouter(routers...)
}
~~~

##以image.NewRouter(opts.daemon.ImageService())为例追查

~~~go
//daemon/daemon.go:1550
func (daemon *Daemon) ImageService() *images.ImageService {
	return daemon.imageService
}

//daemon/daemon.go:86
type Daemon struct {
...
	imageService      *images.ImageService
...
}

//daemon/daemon.go:728
func NewDaemon(ctx context.Context, config *config.Config, pluginStore *plugin.Store) (daemon *Daemon, err error) {
...
	d := &Daemon{
		configStore: config,
		PluginStore: pluginStore,
		startupDone: make(chan struct{}),
	}
...
	d.imageService = images.NewImageService(imgSvcConfig)
...
	return d, nil
}
~~~
