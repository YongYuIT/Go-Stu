# build and install Server.Engine (dockerd)

## download source code
~~~shell script
$ wget https://codeload.github.com/docker/engine/zip/v19.03.13
$ unzip engine-19.03.13.zip
~~~

## build binary file
~~~shell script
$ cd engine-19.03.13/
$ mkdir -p $GOPATH/src/github.com/docker/docker
$ mv ./* $GOPATH/src/github.com/docker/docker/
$ cd $GOPATH/src/github.com/docker/docker/
$ VERSION=testbuild DOCKER_GITCOMMIT=1 ./hack/make.sh binary
can't load package: cannot find module providing package github.com/docker/docker/cmd/dockerd: working directory is not part of a module
$ export GO111MODULE=off
$ export GO15VENDOREXPERIMENT=1
$ VERSION=testbuild DOCKER_GITCOMMIT=1 ./hack/make.sh binary
# github.com/docker/docker/daemon/graphdriver/btrfs
daemon/graphdriver/btrfs/btrfs.go:8:10: fatal error: btrfs/ioctl.h: No such file or directory
$ sudo apt install btrfs-tools
$ VERSION=testbuild DOCKER_GITCOMMIT=1 ./hack/make.sh binary
---> Making bundle: binary (in bundles/binary)
Building: bundles/binary-daemon/dockerd-testbuild
GOOS="" GOARCH="" GOARM=""
Created binary: bundles/binary-daemon/dockerd-testbuild
$ bundles/binary-daemon/dockerd-testbuild --help
$ bundles/binary-daemon/dockerd-testbuild --version
~~~

## install Server.Engine which you build

~~~shell script
$ systemctl cat docker.service
# /lib/systemd/system/docker.service
~~~

* find and replace "ExecStart"
~~~
#ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
ExecStart=/home/yong/Desktop/go_env/go-path/src/github.com/docker/docker/bundles/binary-daemon/dockerd-testbuild -H fd:// --containerd=/run/containerd/containerd.sock
~~~

~~~shell script
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker
$ docker version
Client: Docker Engine - Community
 Version:           19.03.12
 API version:       1.40
 Go version:        go1.13.10
 Git commit:        48a66213fe
 Built:             Mon Jun 22 15:45:44 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server:
 Engine:
  Version:          testbuild
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.14.1
  Git commit:       1
  Built:            Wed Dec 16 09:28:23 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.13
  GitCommit:        7ad184331fa3e55e52b890ea95e65ba581ae3429
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
~~~

* 可见Server.Engine已经改过来了，现在改回去，尝试编译Client

~~~
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
#ExecStart=/home/yong/Desktop/go_env/go-path/src/github.com/docker/docker/bundles/binary-daemon/dockerd-testbuild -H fd:// --containerd=/run/containerd/containerd.sock
~~~

~~~shell script
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker
$  docker version
Client: Docker Engine - Community
 Version:           19.03.12
 API version:       1.40
 Go version:        go1.13.10
 Git commit:        48a66213fe
 Built:             Mon Jun 22 15:45:44 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.12
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.13.10
  Git commit:       48a66213fe
  Built:            Mon Jun 22 15:44:15 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.13
  GitCommit:        7ad184331fa3e55e52b890ea95e65ba581ae3429
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
~~~

# build and install Client (docker client not dockerd)

~~~shell script
$ export GO111MODULE=off
$ export GO15VENDOREXPERIMENT=1
$ VERSION=testbuild DOCKER_GITCOMMIT=1 ./hack/make.sh
No package 'devmapper' found
pkg-config: exit status 1
$ sudo apt-get install libdevmapper-dev
$ VERSION=testbuild DOCKER_GITCOMMIT=1 ./hack/make.sh
~~~

fuck, where is docker and docker-proxy?

## download source code for docker

* https://github.com/docker

~~~shell script
$ wget https://github.com/docker/cli/archive/v20.10.1.zip
$ unzip cli-20.10.1.zip
$ cd $GOPATH/src/github.com/docker/
$ mv cli-20.10.1 .
$ mv cli-20.10.1 cli
$ ls
cli  docker
$ cd cli
$ export GO111MODULE=off
$ export GO15VENDOREXPERIMENT=1
$ export GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct
$ make -f docker.Makefile binary
failed to solve with frontend dockerfile.v0: failed to build LLB: failed to load cache key: failed to do request: Head https://registry.docker-cn.com/v2/library/golang/manifests/1.13.15-alpine: dial tcp 106.14.52.175:443: connect: connection refused
$ make binary
Building statically linked build/docker-linux-amd64
$ build/docker version
Client:
 Version:           unknown-version
 API version:       1.40
 Go version:        go1.14.1
 Git commit:        
 Built:             Wed Dec 23 02:37:18 2020
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          19.03.12
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.13.10
  Git commit:       48a66213fe
  Built:            Mon Jun 22 15:44:15 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.13
  GitCommit:        7ad184331fa3e55e52b890ea95e65ba581ae3429
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
$ build/docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
~~~
