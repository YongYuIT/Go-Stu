package main

import (
	"fmt"
	"hello_hdf/service"
	"os"
)

func main() {
	//所有者，组,其他人
	//rw- --- ---
	dir, err := os.Stat(service.UploadPath)
	isExist := false
	if err == nil {
		isExist = true
	}
	if !isExist || !dir.IsDir() {
		err := os.Mkdir(service.UploadPath, 0700)
		if err != nil {
			fmt.Println("init err-->", err)
			return
		}
	}

	http_service := service.TKHHttpService{Port: 8080}
	http_service.StartService()
}
