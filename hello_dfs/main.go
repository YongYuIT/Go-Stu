package main

import (
	"context"
	"fmt"
	"hello_hdf/service"
	"hello_hdf/tools"
)

func main() {
	ctx := context.Background()
	config, err := tools.GetConfig()
	if err != nil {
		fmt.Println("reac config err-->", err)
		return
	}

	//后续优化代码结构，init工作全部由service对象完成，不直接出现tools包
	err = tools.InitCatchPath(config.GetString("root_cache_path"))
	if err != nil {
		fmt.Println("init cache path err-->", err)
		return
	}

	err = service.InitDb(config.GetString("root_cache_path"))
	if err != nil {
		fmt.Println("init db err-->", err)
		return
	}

	//后续优化代码结构，init工作全部由service对象完成，不直接出现tools包
	err = tools.InitZKRootPath(config.GetString("zk_hosts"))
	if err != nil {
		fmt.Println("init zk err-->", err)
		return
	}

	//这段代码设计上感觉有点问题
	connCtx, unregisterFunc := context.WithCancel(ctx)
	context.WithCancel(ctx)
	go service.RegisterNode(connCtx, config.GetInt("file_api_port"))

	http_service := service.TKHHttpService{Port: config.GetInt("file_api_port")}
	err = http_service.StartService()

	unregisterFunc()
	service.CloseDB()

	if err != nil {
		fmt.Println("start api server err-->", err)
		return
	}
}
