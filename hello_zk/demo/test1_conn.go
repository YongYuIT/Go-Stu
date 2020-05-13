package demo

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

func ConnToZK() {
	conn, _, err := zk.Connect([]string{"0.0.0.0:2181", "0.0.0.0:2182", "0.0.0.0:2183"}, 10*time.Second)
	if err != nil {
		fmt.Println("conn error-->", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn success!")
	//zk.FlagEphemeral 临时节点 临时节点在客户端会话失效后节点自动清除，临时节点下面不能创建子节点 https://www.cnblogs.com/shuiyonglewodezzzzz/p/11208657.html
	//zk.PermAll 允许客户端执行所有操作 https://blog.csdn.net/sdgihshdv/article/details/77660318
	path := "/fuch_index_test"
	conn.Create(path, []byte("init data"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	data, state, err := conn.Get(path)
	if err != nil {
		fmt.Println("get_data_err", err)
		return
	}
	fmt.Println("get init-->", string(data), "-->", state.Version)

	for i := 0; i < 10; i++ {
		state, err = conn.Set(path, []byte("data-->"+strconv.Itoa(i)), state.Version)
		if err != nil {
			fmt.Println("set_data_err", err)
			return
		}
		fmt.Println("set data-->", "-->", state.Version)

		data, state, err = conn.Get(path)
		if err != nil {
			fmt.Println("get_data_err", err)
			return
		}
		fmt.Println("get data-->", string(data), "-->", state.Version)
	}
	return
}
