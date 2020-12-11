# modify to go mod

* Settings > Go > Go modules > Enable Go modules integration > Environment > GOPROXY=https://goproxy.cn,direct
* go mod init
* rm -rf vendor
* go get -u

# gen go interface files by tars file

~~~shell script
$ tars2go Servant.tars
~~~

# write impls and run

* config "--config=config.conf" to start params:
* run server proj

# test with client

2020-12-11 00:28:58.308|application.go:initConfig:73|ERROR|server config path empty
0 3702

至此，TestHelloServer服务与tars framework没有任何关系，下步发布服务到tars framework

# 发布服务

ref to: https://github.com/TarsCloud/TarsGo/blob/master/docs/tars_go_quickstart.md#%E6%9C%8D%E5%8A%A1%E9%83%A8%E7%BD%B2

~~~shell script
$ make && make tar
make: *** No targets specified and no makefile found.  Stop.
$ sudo apt-get update
$ sudo apt-get install gcc build-essential
$ make && make tar
make: *** No targets specified and no makefile found.  Stop.
~~~

突然发现TarsGo/tars/tools/create_tars_server_gomod.sh，应该用这个，fuck failed !!!
