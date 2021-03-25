package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 5000, "remote ssh port")
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("start server err-->", err)
		return
	}
	defer listen.Close()

	for {
		//阻塞等待客户端
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept new conn failed-->", err)
			return
		}
		//创建协程
		go Handler(conn)
	}
}
