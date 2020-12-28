* stu of Unix域套接字

在Linux系统中，有很多进程间通信方式，套接字（Socket）就是其中的一种。但传统的套接字的用法都是基于TCP/IP协议栈的，需要指定IP地址。如果不同主机上的两个进程进行通信，当然这样做没什么问题。但是，如果只需要在一台机器上的两个不同进程间通信，还要用到IP地址就有点大材小用了。

其实很多人并不一定知道，对于套接字来说，还存在一种叫做Unix域套接字的类别，专门用来解决这个问题。其API的掉用方法基本上和普通TCP/IP的套接字一样，只是有些许差别。

# network type

~~~go
func ResolveUnixAddr(network, address string) (*UnixAddr, error) {
	switch network {
	case "unix", "unixgram", "unixpacket":
		return &UnixAddr{Name: address, Net: network}, nil
	default:
		return nil, UnknownNetworkError(network)
	}
}
~~~

* "unix" --> SOCK_STREAM --> TCP
* "unixgram" --> SOCK_DGRAM --> UDP
* "unixpacket" --> SOCK_SEQPACKET --> ???