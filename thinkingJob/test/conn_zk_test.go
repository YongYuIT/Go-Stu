package test

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/go-zookeeper/zk"

func Test_connToZK(test *testing.T) {
	//start zk with docker
	//docker run --name zookeeper-1 --restart always -d -p 2181:2181 zookeeper
	//docker exec -it zookeeper-1 sh
	//zkCli.sh
	//ls /
	//create /zk-book 123
	//ls /zk-book
	//get /zk-book
	//set /zk-book 321
	//get /zk-book
	zkList := []string{"0.0.0.0:2181"}
	conn, env, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println("conn to zk error-->", err)
		return
	}
	for {
		envMsg := <-env
		fmt.Println("print env--------------------------------------------------------------------------------------start")
		fmt.Println(envMsg.Type, "-->", envMsg.State, "-->", envMsg.Server, "-->", envMsg.Path, "-->", envMsg.Err)
		fmt.Println("print env--------------------------------------------------------------------------------------end")
		fmt.Println("conn to zk-->", conn.State())
	}
}
