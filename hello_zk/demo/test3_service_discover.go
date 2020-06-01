package demo

import (
	"bufio"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
	"time"
)

//zk临时节点的生命周期和客户端会话绑定在一起，客户端会话失效，则这个节点就会被自动清除
//这一点很适合用来做服务发现

const root_path string = "/test_conn_root"

func MainTest() {
	err := initRootPath()
	if err != nil {
		fmt.Println("init root path failed")
		return
	}

	msg_buss := make(map[string]chan string)
	inputReader := bufio.NewReader(os.Stdin)
	for {
		str, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("input err-->", err)
			continue
		}
		str = str[:len(str)-1]
		if strings.EqualFold("STOP", str) {
			return
		}
		fmt.Println("handle-->", str)
		if strings.Contains(str, "CREATE") {
			//like "CREATE qaz123"
			cid := strings.Split(str, " ")[1]
			msg_buss[cid] = make(chan string)
			if msg_buss[cid] == nil {
				fmt.Println("client has exit!")
			}
			go CreateClient(cid, msg_buss[cid])
		}

		if strings.Contains(str, "SEND") {
			//like SEND qaz123 helloword
			inputs := strings.Split(str, " ")
			cid := inputs[1]
			msg := inputs[2]
			msg_buss[cid] <- msg
		}
	}
}

func initRootPath() error {
	conn := createConn()
	if conn != nil {
		fmt.Println("conn zk success!")
		defer conn.Close()
	} else {
		fmt.Println("conn zk failed!")
		return fmt.Errorf("conn zk failed")
	}

	exist, state, err := conn.Exists(root_path)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return fmt.Errorf("read zk info falied")
	}
	if !exist {
		//创建永久节点
		result_str, err := conn.Create(root_path, []byte("init data"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("root path create failed!")
			return fmt.Errorf("create root path failed")
		} else {
			fmt.Println("root path create success-->", result_str)
		}
	}
	return nil
}

//创建新客户端（服务节点）
func CreateClient(cid string, msg_bus <-chan string) {
	conn := createConn()
	if conn != nil {
		fmt.Println("conn zk success!")
		defer conn.Close()
	} else {
		fmt.Println("conn zk failed!")
		return
	}

	sub_path := root_path + "/" + cid
	exist, state, err := conn.Exists(sub_path)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return
	}
	if !exist {
		//创建临时节点
		result_str, err := conn.Create(sub_path, []byte("init data"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("create path failed-->", err)
			return
		} else {
			fmt.Println("root path create success-->", result_str)
		}
	} else {
		fmt.Println("client exist-->%s", cid)
	}

	for {
		msg := <-msg_bus
		if strings.EqualFold(cid+"_stop", msg) {
			//like SEND qaz123 qaz123_stop
			break
		} else {
			fmt.Println("recv msg-->", msg)
		}
	}
}

func createConn() *zk.Conn {
	conn, _, err := zk.Connect([]string{"0.0.0.0:2181", "0.0.0.0:2182", "0.0.0.0:2183"}, 10*time.Second)
	if err != nil {
		fmt.Println("conn error-->", err)
		return nil
	}
	return conn
}
