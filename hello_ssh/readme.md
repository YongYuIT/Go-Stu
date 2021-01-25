# build target env
~~~shell script
docker run -it --name test_env -p 5022:22 ubuntu bash
docker cp sources.list test_env:/etc/apt/sources.list
~~~
in container
~~~shell script
apt-get update
apt-get install openssl ssh net-tools
service ssh start
adduser yong
ifconfig
#get ip of container: 192.168.100.3
ssh yong@192.168.100.3
~~~



~~~shell script
ssh-copy-id -i ~/.ssh/id_rsa.pub yong@192.168.100.3
ssh yong@192.168.100.3
~~~

# usage

~~~shell script
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=
go build
####windows
.\hello_ssh.exe -D /home/yong/tttmp -S test_dir -h 192.168.186.140 -i .\id_rsa -p 22 -t upload -u yong
####linux
./hello_ssh -h 192.168.100.3 -p 22 -i /home/yong/.ssh/id_rsa -u yong -S ../ -D /home/yong/tmp1
~~~