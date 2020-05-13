package demo

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

func CreatePaths() {
	conn, _, err := zk.Connect([]string{"0.0.0.0:2181", "0.0.0.0:2182", "0.0.0.0:2183"}, 10*time.Second)
	if err != nil {
		fmt.Println("conn error-->", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn success!")
	path := "/fuch_index_test"
	exist, state, err := conn.Exists(path)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return
	}
	if !exist {
		conn.Create(path, []byte("init data"), 0, zk.WorldACL(zk.PermAll))
	}

	for i := 0; i < 4; i++ {
		sub_path := path + "/index_" + strconv.Itoa(i)
		exist, state, err = conn.Exists(sub_path)
		if err != nil {
			fmt.Println("read path state err-->", err, state.Version)
			return
		}
		if !exist {
			result, err := conn.Create(sub_path, []byte("init data "+strconv.Itoa(i)), 0, zk.WorldACL(zk.PermAll))
			if err != nil {
				fmt.Println("create subpath err-->", err, result)
				return
			}
		}
		data, state, err := conn.Get(sub_path)
		if err != nil {
			fmt.Println("get_data_err", err)
			return
		}
		fmt.Println("get init-->", string(data), "-->", state.Version)
	}
}
