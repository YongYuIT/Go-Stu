## docker network type

Docker 提供三种网络模式

* none : https://www.cnblogs.com/sunqingliang/p/12748691.html
* host : https://www.cnblogs.com/sunqingliang/p/12748691.html
* bridge : https://www.cnblogs.com/sunqingliang/p/12741387.html

Docker 提供三种 user-defined 网络驱动

* bridge
* overlay
* macvlan

overlay 和 macvlan 用于创建跨主机的网络

## analyse docker network bridge creating process

~~~go
//daemon/daemon_unix.go:848
func (daemon *Daemon) initNetworkController(config *config.Config, activeSandboxes map[string]interface{}) (libnetwork.NetworkController, error) {
...
	if n, _ := controller.NetworkByName("none"); n == nil {
...
	if n, _ := controller.NetworkByName("host"); n == nil {
...
	if n, err := controller.NetworkByName("bridge"); err == nil {
...
        //创建默认网桥，docker0（用ifconfig可以看到），配置在/etc/docker/daemon.json-->bip
		if err := initBridgeDriver(controller, config); err != nil {
...
}

//daemon/daemon_unix.go:927
func initBridgeDriver(controller libnetwork.NetworkController, config *config.Config) error {
...
	_, err = controller.NewNetwork("bridge", "bridge", "",
		libnetwork.NetworkOptionEnableIPv6(config.BridgeConfig.EnableIPv6),
		libnetwork.NetworkOptionDriverOpts(netOption),
		libnetwork.NetworkOptionIpam("default", "", v4Conf, v6Conf, nil),
		libnetwork.NetworkOptionDeferIPv6Alloc(deferIPv6Alloc))
...
}

//vendor/github.com/docker/libnetwork/controller.go:709
func (c *controller) NewNetwork(networkType, name string, id string, options ...NetworkOption) (Network, error) {
...
	err = c.addNetwork(network)
...
}

//vendor/github.com/docker/libnetwork/controller.go:1004
func (c *controller) addNetwork(n *network) error {
...
	if err := d.CreateNetwork(n.id, n.generic, n, n.getIPData(4), n.getIPData(6)); err != nil {
...
}

//vendor/github.com/docker/libnetwork/drivers/bridge/bridge.go:609
func (d *driver) CreateNetwork(id string, option map[string]interface{}, nInfo driverapi.NetworkInfo, ipV4Data, ipV6Data []driverapi.IPAMData) error {
...
//这段代码主要描述了bridge创建过程，主要步骤在bridgeSetup队列里面，通过bridgeSetup.queueStep添加
//里面的for循环，批量添加（添加之前判断条件）步骤，很有特色
	// Prepare the bridge setup configuration
	bridgeSetup := newBridgeSetup(config, bridgeIface)

	// If the bridge interface doesn't exist, we need to start the setup steps
	// by creating a new device and assigning it an IPv4 address.
	bridgeAlreadyExists := bridgeIface.exists()
	if !bridgeAlreadyExists {
		bridgeSetup.queueStep(setupDevice)
		bridgeSetup.queueStep(setupDefaultSysctl)
	}

	// For the default bridge, set expected sysctls
	if config.DefaultBridge {
		bridgeSetup.queueStep(setupDefaultSysctl)
	}

	// Even if a bridge exists try to setup IPv4.
	bridgeSetup.queueStep(setupBridgeIPv4)

	enableIPv6Forwarding := d.config.EnableIPForwarding && config.AddressIPv6 != nil

	// Conditionally queue setup steps depending on configuration values.
	for _, step := range []struct {
		Condition bool
		Fn        setupStep
	}{
		// Enable IPv6 on the bridge if required. We do this even for a
		// previously  existing bridge, as it may be here from a previous
		// installation where IPv6 wasn't supported yet and needs to be
		// assigned an IPv6 link-local address.
		{config.EnableIPv6, setupBridgeIPv6},

		// We ensure that the bridge has the expectedIPv4 and IPv6 addresses in
		// the case of a previously existing device.
		{bridgeAlreadyExists && !config.InhibitIPv4, setupVerifyAndReconcile},

		// Enable IPv6 Forwarding
		{enableIPv6Forwarding, setupIPv6Forwarding},

		// Setup Loopback Addresses Routing
		{!d.config.EnableUserlandProxy, setupLoopbackAddressesRouting},

		// Setup IPTables.
		{d.config.EnableIPTables, network.setupIP4Tables},

		// Setup IP6Tables.
		{config.EnableIPv6 && d.config.EnableIP6Tables, network.setupIP6Tables},

		//We want to track firewalld configuration so that
		//if it is started/reloaded, the rules can be applied correctly
		{d.config.EnableIPTables, network.setupFirewalld},
		// same for IPv6
		{config.EnableIPv6 && d.config.EnableIP6Tables, network.setupFirewalld6},

		// Setup DefaultGatewayIPv4
		{config.DefaultGatewayIPv4 != nil, setupGatewayIPv4},

		// Setup DefaultGatewayIPv6
		{config.DefaultGatewayIPv6 != nil, setupGatewayIPv6},

		// Add inter-network communication rules.
		{d.config.EnableIPTables, setupNetworkIsolationRules},

		//Configure bridge networking filtering if ICC is off and IP tables are enabled
		{!config.EnableICC && d.config.EnableIPTables, setupBridgeNetFiltering},
	} {
		if step.Condition {
			bridgeSetup.queueStep(step.Fn)
		}
	}

	// Apply the prepared list of steps, and abort at the first error.
	bridgeSetup.queueStep(setupDeviceUp)
	return bridgeSetup.apply()
}
~~~
### 以bridgeSetup.queueStep(setupDevice)为例，追溯系统调用过程

~~~go
//vendor/github.com/docker/libnetwork/drivers/bridge/setup_device.go:15
func setupDevice(config *networkConfiguration, i *bridgeInterface) error {
...
		return ioctlCreateBridge(config.BridgeName, hwAddr.String())
}

//vendor/github.com/docker/libnetwork/drivers/bridge/netlink_deprecated_linux.go:107
func ioctlCreateBridge(name, macAddr string) error {
...
	nameBytePtr, err := syscall.BytePtrFromString(name)
	if err != nil {
		return err
	}
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(s), ioctlBrAdd, uintptr(unsafe.Pointer(nameBytePtr))); err != 0 {
		return err
	}
	return ioctlSetMacAddress(name, macAddr)
}
~~~
## golang syscall
详见 https://github.com/YongYuIT/Go-Stu/tree/master/hello-docker/try_2_syscall
## linux ioctl 创建网桥
详见 linux_bradge.md
## linux netlink （待续）
