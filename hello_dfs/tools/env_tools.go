package tools

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
)

var CatchPath = "." + UploadPath

const UploadPath = "/th_dfs_path/"
const ZKRootPath = "/tk_dfs_nodes"

func InitCatchPath(path_start string) error {
	CatchPath := path_start + UploadPath
	dir, err := os.Stat(CatchPath)
	if err == nil && dir.IsDir() {
		return nil
	}
	//所有者，组,其他人
	//rwx --- ---
	err = os.Mkdir(CatchPath, 0700)
	if err != nil {
		fmt.Println("init err-->", err)
		return err
	}
	return nil
}

var ZKHOSES []string = nil

func InitZKRootPath(zk_hosts string) error {
	ZKHOSES = strings.Split(zk_hosts, ",")
	conn := CreateZKConn(ZKHOSES)
	if conn != nil {
		fmt.Println("conn zk success!")
		defer conn.Close()
	} else {
		fmt.Println("conn zk failed!")
		return fmt.Errorf("conn zk failed")
	}

	exist, state, err := conn.Exists(ZKRootPath)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return fmt.Errorf("read zk info falied")
	}
	if !exist {
		//创建永久节点
		result_str, err := conn.Create(ZKRootPath, []byte("init data"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("root path create failed!")
			return fmt.Errorf("create root path failed")
		} else {
			fmt.Println("root path create success-->", result_str, ZKRootPath)
		}
	}
	return nil
}

func getIpPort() string {
	ip, err := externalIP()
	if err != nil {
		fmt.Println("get ip error")
		return ""
	}
	return ip.String()
}
