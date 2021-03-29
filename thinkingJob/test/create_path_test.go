package test

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"testing"
	"time"
)

/*
ZK里面有四种类型的节点：

永久 0
永久且递增 2 zk.FlagSequence
临时 1 zk.FlagEphemeral
临时且递增 3

永久节点需要手动删除，临时节点会话中断即删除
*/

func getNewConn() *zk.Conn {
	zkList := []string{"0.0.0.0:2181"}
	conn, env, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println("conn to zk error-->", err)
		return nil
	}
	for {
		select {
		case msg := <-env:
			if msg.State == zk.StateHasSession {
				goto exit
			}
		case <-time.After(15 * time.Second):
			return nil
		}
	}
exit:
	return conn
}

func Test_createPermanentPath(test *testing.T) {
	conn := getNewConn()
	if conn == nil {
		return
	}
	//zk.WorldACL(zk.PermAll) 设置访问模式
	conn.Create("/test_createPermanentPath", []byte("test1"), 0, zk.WorldACL(zk.PermAll))
	conn.Close()

	//docker exec -it zookeeper-1 sh
	//zkCli.sh
	//ls /
}
