ref to https://tarscloud.gitbook.io/tarsdocs/rumen/installation/docker
# build network env
~~~shell script
# 创建一个名为tars的桥接(bridge)虚拟网络，网关172.25.0.1，网段为172.25.0.0
$ docker network create -d bridge --subnet=172.25.0.0/16 --gateway=172.25.0.1 tars
4019bea1c2a3d737ea3bd8234014ebc3f0097ea0ad4dd93bf5f604ae0c4e39fe

$ ifconfig
br-4019bea1c2a3: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.25.0.1  netmask 255.255.0.0  broadcast 172.25.255.255
        ether 02:42:79:50:e8:14  txqueuelen 0  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
~~~
# build mysql env
~~~shell script
$ docker images | grep mysql
mysql                          5.7.32              ae0658fdbad5        2 weeks ago         449MB
mysql                          latest              990386cbd5c0        19 months ago       443MB
$ docker run -d \
    --net=tars \
    -e MYSQL_ROOT_PASSWORD="123456" \
    --ip="172.25.0.2" \
    -v /data/framework-mysql:/var/lib/mysql \
    -v /etc/localtime:/etc/localtime \
    --name=tars-mysql \
    mysql:5.7.32
~~~
# build tars framework
~~~shell script
$ docker pull tarscloud/framework:stable
# 挂载的/etc/localtime是用来设置容器时区的，若没有可以去掉
  # 3000端口为web程序端口
  # 3001端口为web授权相关服务端口
$ docker run -d \
      --name=tars-framework \
      --net=tars \
      -e MYSQL_HOST="172.25.0.2" \
      -e MYSQL_ROOT_PASSWORD="123456" \
      -e MYSQL_USER=root \
      -e MYSQL_PORT=3306 \
      -e REBUILD=false \
      -e INET=eth0 \
      -e SLAVE=false \
      --ip="172.25.0.3" \
      -v /data/framework:/data/tars \
      -v /etc/localtime:/etc/localtime \
      -p 3000:3000 \
      -p 3001:3001 \
      tarscloud/framework:stable
~~~
http://0.0.0.0:3000/ 搞定