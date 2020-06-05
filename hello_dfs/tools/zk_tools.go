package tools

import (
	"context"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

func CreateZKConn(zk_host []string) *zk.Conn {
	conn, _, err := zk.Connect(zk_host, 10*time.Second)
	if err != nil {
		fmt.Println("conn error-->", err)
		return nil
	}
	return conn
}

func CreateNodeInfo(node_name string, ctx context.Context, port int) error {
	conn := CreateZKConn(ZKHOSES)
	if conn != nil {
		fmt.Println("conn zk success!")
		defer conn.Close()
	} else {
		fmt.Println("conn zk failed!")
		return fmt.Errorf("conn zk failed")
	}

	sub_path := ZKRootPath + "/" + node_name
	exist, state, err := conn.Exists(sub_path)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return err
	}
	if !exist {
		//创建临时节点
		result_str, err := conn.Create(sub_path, []byte(getIpPort()+":"+strconv.Itoa(port)), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("create path failed-->", err)
			return err
		} else {
			fmt.Println("node path create success-->", result_str)
		}
	} else {
		fmt.Println("client exist-->%s", node_name)
	}
	msg := <-ctx.Done()
	fmt.Println("zk conn end-->", msg)
	return nil
}
