# gen new proj
~~~shell script
$ bash ~/codes/TarsGo/tars/tools/create_tars_server_gomod.sh TestHelloApp1 TestHelloSer TestHelloSvan TestHelloMod
~~~

# gen go interface files by tars file

~~~shell script
$ tars2go Servant.tars
~~~

# write impls and run

* config "--config=config.conf" to start params:
* run server proj

# test with client

2020-12-11 02:04:35.451|application.go:initConfig:73|ERROR|server config path empty
0 37035

# make and tar
~~~shell script
$ make && make tar
make: *** No targets specified and no makefile found.  Stop.
~~~
replace makefile "$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/TarsCloud/TarsGo/tars/makefile.tars.gomod))"

with "$(foreach path,$(libpath),$(eval -include /home/yong/codes/TarsGo/tars/makefile.tars.gomod))"

# 发布服务

ref to: https://github.com/TarsCloud/TarsGo/blob/master/docs/tars_go_quickstart.md#%E6%9C%8D%E5%8A%A1%E9%83%A8%E7%BD%B2

> servant obj name not found:TestHelloApp1.TestHelloSer.TestHelloSvanObj