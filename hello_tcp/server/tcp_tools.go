package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Handler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	//读取客户端发送的内容
	//check client
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read from server err1-->", err)
		return
	}
	checkStr := string(buf[:n])
	if !strings.EqualFold(getCheckStr(), checkStr) {
		fmt.Println("cannot confirm client")
		return
	} else {
		conn.Write([]byte("checked!"))
	}

	//get fileName
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("get file name err-->", err)
		return
	}
	fileName := string(buf[:n])
	//获取客户端ip+port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": 客户端传输的文件名为--" + fileName)
	conn.Write([]byte("ready!"))

	err = os.Mkdir("./"+checkStr, os.ModePerm)
	if err != nil {
		fmt.Println("create file err-->", err)
		return
	}
	file, err := os.Create("./" + checkStr + "/" + fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("create file err-->", err)
		return
	}
	//循环接收客户端传递的文件内容
	finishTag := "####finish####"
	fmt.Println("start recv-->", file.Name())
	for {
		n, _ := conn.Read(buf)
		if strings.Contains(string(buf[:n]), finishTag) {
			fmt.Println("upload file-->", file.Name(), "-->", string(buf[:n]))
			conn.Write([]byte("uploaded!"))
			if n > len(finishTag) {
				file.Write(buf[:n-len(finishTag)])
			}
			break
		}
		file.Write(buf[:n])
	}
	fmt.Println("end recv-->", file.Name())
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("get cmd err-->", err)
		return
	}
	cmdStr := string(buf[:n])
	fmt.Println("cmd: ", cmdStr)
	cmd := exec.Command("bash", "-c", cmdStr)
	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println("exec cmd err-->", err)
		return
	}
	conn.Write(cmdOut)

	fmt.Println(addr + ": 协程结束")
	runtime.Goexit()
}

func getCheckStr() string {
	t1 := time.Now().Year()   //年
	t2 := time.Now().Month()  //月
	t3 := time.Now().Day()    //日
	t4 := time.Now().Hour()   //小时
	t5 := time.Now().Minute() //分钟
	date := fmt.Sprintf("%d-%d-%d %d:%d", t1, t2, t3, t4, t5)
	key := md5.Sum([]byte(date + "-hello-webase"))
	return hex.EncodeToString(key[:])
}
