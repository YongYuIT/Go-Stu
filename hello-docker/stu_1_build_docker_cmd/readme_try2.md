# build and install all parts of docker

## download source code
~~~shell script
$ wget https://github.com/moby/moby/archive/v20.10.1.zip
$ unzip v20.10.1.zip
~~~

## build proj
~~~shell script
$ cd $GOPATH/src/github.com/docker
$ mv docker docker_bak_19.03.13
$ mv moby-20.10.1 ./
$ mv moby-20.10.1 docker
$ cd docker
$ export GO111MODULE=off
$ export GO15VENDOREXPERIMENT=1
$ DOCKER_GITCOMMIT=1 hack/make.sh binary
~~~

make binary: build statically linked linux binaries

make dynbinary: build dynamically linked linux binaries
