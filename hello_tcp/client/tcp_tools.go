package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func tcpHandle(serverhost string, port int, zipfile string, cmd string) error {
	buf := make([]byte, 2048)
	//获取文件信息
	fileInfo, err := os.Stat(zipfile)
	if err != nil {
		fmt.Println("get file err-->", err)
		return err
	}
	//创建客户端连接
	conn, err := net.Dial("tcp", serverhost+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	//check client
	key := getCheckStr()
	conn.Write([]byte(key))
	respLength, err := conn.Read(buf)
	resp := string(buf[:respLength])
	if resp == "checked!" {
		fmt.Println("client checked!")
	} else {
		return fmt.Errorf("uncheckd client")
	}

	//文件名称
	fileName := fileInfo.Name()
	//文件大小
	fileSize := fileInfo.Size()
	//发送文件名称到服务端
	conn.Write([]byte(fileName))
	respLength, err = conn.Read(buf)
	resp = string(buf[:respLength])
	if resp == "ready!" {
		//发送文件数据
		SendFile(zipfile, fileSize, conn)
	}
	fmt.Println("wait for uploaded")
	respLength, err = conn.Read(buf)
	resp = string(buf[:respLength])
	if resp == "uploaded!" {

		cmd = strings.Replace(cmd, "KEY_VAL", key, -1)
		cmd = strings.Replace(cmd, "ZIP_VAL", fileName, -1)
		cmd = strings.Replace(cmd, "PROJ_VAL", projMap[projNum], -1)
		cmd = strings.Replace(cmd, "VER_VAL", varName, -1)

		conn.Write([]byte(cmd))
	} else {
		fmt.Println("upload err")
	}

	respLength, err = conn.Read(buf)
	resp = string(buf[:respLength])
	fmt.Println("exec-->", resp)

	return nil
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

//发送文件到服务端
func SendFile(filePath string, fileSize int64, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var count int64
	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if io.EOF == err {
			fmt.Println("文件传输完成 EOF-->", err, "-->", n)
			//告诉服务端结束文件接收
			conn.Write([]byte("####finish####"))
			return
		} else {
			count += int64(n)
			sendPercent := float64(count) / float64(fileSize) * 100
			value := fmt.Sprintf("%.2f", sendPercent)
			//打印上传进度
			fmt.Println("文件上传："+value+"%-->", count, "/", fileSize)
			//发送给服务端
			conn.Write(buf[:n])
		}
	}
}
