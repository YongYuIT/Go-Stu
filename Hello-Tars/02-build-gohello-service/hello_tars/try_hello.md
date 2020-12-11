ref to https://github.com/TarsCloud/TarsGo/blob/master/README.zh.md
# download dep and install tools
~~~shell script
$ go env -w GOPROXY=https://goproxy.cn 
$ go env -w GO111MODULE=auto
$ go get -u github.com/TarsCloud/TarsGo/tars
go: found github.com/TarsCloud/TarsGo/tars in github.com/TarsCloud/TarsGo v1.1.5
$ go install $GOPATH/src/github.com/TarsCloud/TarsGo/tars/tools/tars2go
can't load package: package /home/yong/Desktop/go_env/go-path/src/github.com/TarsCloud/TarsGo/tars/tools/tars2go: cannot find package
~~~
> fix "can't load package: package /home/yong/Desktop/go_env/go-path/src/github.com/TarsCloud/TarsGo/tars/tools/tars2go: cannot find package"

~~~shell script
$ cd ~/codes
$ git clone https://github.com/TarsCloud/TarsGo.git
$ cd TarsGo/tars/tools/tars2go/
$ go build && go install
~~~
# gen go interface files by tars file
~~~shell script
$ tars2go --outdir=./vendor hello.tars
$ rm -rf vendor
~~~
# write impls and run

ref to: https://github.com/TarsCloud/TarsGo/blob/master/README.zh.md#13-%E6%8E%A5%E5%8F%A3%E5%AE%9E%E7%8E%B0

> 2020-12-10 19:35:15.887|application.go:initConfig:73|ERROR|server config path empty
  panic: runtime error: invalid memory address or nil pointer dereference

create config.conf ref to: https://github.com/TarsCloud/TarsGo/blob/master/README.zh.md#14-%E6%9C%8D%E5%8A%A1%E7%AB%AF%E9%85%8D%E7%BD%AE

config "--config=config.conf" to start params:
  * Edit configurations... > Program arguments
  
> panic: empty servant

fuck failed ,try an other way

ref to: https://github.com/TarsCloud/TarsGo/blob/master/docs/tars_go_quickstart.md

~~~shell script
$ bash ~/codes/TarsGo/tars/tools/create_tars_server.sh TestHelloApp TestHelloServer TestHelloServant
$ mv $GOPATH/src/TestHelloApp Go-Stu/Hello-Tars/02-build-gohello-service/
~~~
 bye!!!