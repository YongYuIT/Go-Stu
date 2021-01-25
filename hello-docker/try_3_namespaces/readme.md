# desc

Linux Namespace 是 Kernel 的 一个功能,它可以隔离 一 系列的系统资源

# Linux 一共实现了 6 种不同类型的 Namespace

|类型|系统调用参数|内核版本|隔离系统资源|
|----|----|----|----|
|Mount Namespace|CLONE_NEWNS|2.4.19|隔离各个进程看到的挂载点视图|
|UTS Namespace|CLONE_NEWUTS|2.6.19|用来隔离nodename和domainname两个系统标识。在UTS Namespace里面,每个Namespace允许有自己的hostname|
|IPC Namespace|CLONE_NEWIPC|2.6.19|隔离System V IPC和POSIX message queues，即主要用来隔离消息队列|
|PID Namespace|CLONE_NEWPID|2.6.24|隔离进程ID，同一进程在不同的PID Namespace里可以拥有不同的PID|
|Network Namespace|CLONE_NEWNET|2.6.29|隔离网络设备、 IP地址端口等网络栈，可以让每个容器拥有自己独立的(虚拟的)网络设备|
|User Namespace|CLONE_NEWUSER|3.8|隔离用户的用户组ID，同一进程的User ID和Group ID在User Namespace内外可以是不同的|