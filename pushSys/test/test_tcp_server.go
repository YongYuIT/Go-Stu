package test

import (
	"fmt"
	"net"
	"time"
)

import "../tool"

func TestOfTcpServer() {
	ip, err := tool.GetIp()
	if err != nil {
		fmt.Println("bind error -> " + err.Error())
		return
	}
	fmt.Println("start main -> " + ip)
	port := ":6666"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println("bind error -> " + err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("listen error -> " + err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error -> " + err.Error())
			continue
		}
		fmt.Println(conn.RemoteAddr().String() + " join !");
		conn.Write([]byte("this is reply -> " + time.Now().String()))
		conn.Close()
	}
}
