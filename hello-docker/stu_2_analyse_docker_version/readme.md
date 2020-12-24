# where is the func
* exec cmd: docker version --help
* search str "Show the Docker version information" in source code
~~~go
	cmd := &cobra.Command{
		Use:   "version [OPTIONS]",
		Short: "Show the Docker version information",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(dockerCli, &opts)
		},
	}
~~~
the func is: runVersion(dockerCli, &opts)

~~~go
//cli/command/system/version.go:137    func runVersion
sv, err := dockerCli.Client().ServerVersion(context.Background())
~~~

~~~go
//vendor/github.com/docker/docker/client/version.go:11
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error) {
	resp, err := cli.get(ctx, "/version", nil, nil)
...
}
~~~