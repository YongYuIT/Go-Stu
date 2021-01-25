# desc
* 网桥就是现实世界中的交换机，根据mac地址（内部维护了mac地址和交换机端口的对应关系）实现数据包转发，运行在第二层
# create with brctl
~~~shell script
$ sudo apt-get install bridge-utils
$ sudo apt install uml-utilities
####创建虚拟网卡，USER拥有创建网卡权限
$ sudo tunctl -t tap0 -u ${USER}
Set 'tap0' persistent and owned by uid 1000
$ sudo lshw -C network
  *-network                 
       description: Ethernet interface
       product: 82545EM Gigabit Ethernet Controller (Copper)
       vendor: Intel Corporation
       physical id: 1
       bus info: pci@0000:02:01.0
       logical name: ens33
       version: 01
       serial: 00:0c:29:0f:e8:a2
       size: 1Gbit/s
       capacity: 1Gbit/s
       width: 64 bits
       clock: 66MHz
       capabilities: pm pcix bus_master cap_list rom ethernet physical logical tp 10bt 10bt-fd 100bt 100bt-fd 1000bt-fd autonegotiation
       configuration: autonegotiation=on broadcast=yes driver=e1000 driverversion=7.3.21-k8-NAPI duplex=full ip=192.168.186.140 latency=0 link=yes mingnt=255 multicast=yes port=twisted pair speed=1Gbit/s
       resources: irq:19 memory:fd5c0000-fd5dffff memory:fdff0000-fdffffff ioport:2000(size=64) memory:fd500000-fd50ffff
  *-network:0
       description: Ethernet interface
       physical id: 2
       logical name: docker0
       serial: 02:42:d3:f6:81:14
       capabilities: ethernet physical
       configuration: broadcast=yes driver=bridge driverversion=2.3 firmware=N/A ip=192.168.100.1 link=no multicast=yes
  *-network:1 DISABLED
       description: Ethernet interface
       physical id: 3
       logical name: tap0
       serial: 56:f8:c7:2a:5e:a7
       size: 10Mbit/s
       capabilities: ethernet physical
       configuration: autonegotiation=off broadcast=yes driver=tun driverversion=1.6 duplex=full link=no multicast=yes port=twisted pair speed=10Mbit/s
####创建网桥，并将实体网卡与虚拟网卡都连接到网桥上
$ sudo brctl addbr br0 #创建网桥
$ sudo brctl addif br0 ens33  #连接实体网卡
$ sudo brctl addif br0 tap0  #连接虚拟网卡
####添加实体网卡会导致VM虚拟机里面的网络用不了，原因不明，断开命令如下
####sudo brctl delif br0 ens33
####sudo brctl show
####设置网桥具体信息，并启动
$ sudo ifconfig br0 192.168.101.1 netmask 255.255.255.0 up #网桥ip
$ route #查询本机当前路由表：Destination目标网段或者主机；Gateway网关地址；Iface输出接口，即网卡；
$ sudo route add default gw 192.168.101.1  #路由网关
####设置“路由网关”之后会导致VM虚拟机里面的网络用不了，原因不明，清除命令如下
####sudo route del default gw 192.168.101.1
~~~
