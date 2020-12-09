# PLAN1: run go source file in docker

## step1: download fabric-sample proj

git clone fabric-sample

cd test-network/

## setp2: run sample network

#### config fabric-samples/bin and start test-network

~~~shell script
$ bash network.sh up createChannel -i 2.0.0
~~~

## step3: modify sample network 

##### stop peer0.org1.example.com container

~~~shell script
$ docker stop peer0.org1.example.com
$ docker rm peer0.org1.example.com
~~~

####copy peer0.org1.example.com docker-compose file and modify,


##### add volumes:

* /yong/codes/fabric:/yong/codes/fabric
* $GOPATH:/go-path
* $GOROOT:/go
* /yong/codes/Go-Stu/debug_with_docker_fabric/start.sh:/yong/cmd/start.sh

##### add environment:

* GOPATH=/go-path
* GOROOT=/go

##### modify working_dir:

* /yong/cmd

##### modify command:

* sh start.sh

## step4: restart and test

#### restart peer0.org1.example.com container

~~~shell script
$ docker-compose -f net/peer0.org1.example.com.yaml up -d
~~~

fuck go not found!!!

#### test with chain code

~~~shell script
$ bash network.sh deployCC
~~~

# PLAN2: run peer local(not in docker)

## step2: gen config files



## step1: set env

~~~shell script
$ docker history hyperledger/fabric-peer:2.0.0
<missing>           9 months ago        /bin/sh -c #(nop)  ENV FABRIC_CFG_PATH=/etc/…   0B
$ docker exec -it peer0.org2.example.com sh
# echo $FABRIC_CFG_PATH
# ls $FABRIC_CFG_PATH
# cat $FABRIC_CFG_PATH/core.yaml
~~~

