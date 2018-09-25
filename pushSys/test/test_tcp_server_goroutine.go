package test

import (
	"fmt"
	"net"
	"time"
)

import (
	"../tool"
	"strings"
	"reflect"
)

func TestOfTcpServerGoroutine() {
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
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	conn.Write([]byte("conned -> " + time.Now().String()))
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	inputCache := make([]byte, 128)
	fmt.Println(conn.RemoteAddr().String() + " waiting...")
	for {
		read_len, err := conn.Read(inputCache);
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + " error: " + err.Error() + " -> " + reflect.TypeOf(err).String())
			op_error, ok := err.(*net.OpError)
			if ok && op_error.Timeout() {
				fmt.Println(conn.RemoteAddr().String() + " timeout !")
				break
			}
			continue
		}
		if read_len == 0 {
			fmt.Println(conn.RemoteAddr().String() + " no input stop !")
			break //stop by client
		}
		input_str := strings.TrimSpace(string(inputCache[:read_len]))
		if input_str == "stop" {
			fmt.Println(conn.RemoteAddr().String() + " stop !")
			break
		}
		conn.Write([]byte("reply for -> " + input_str))
		inputCache = make([]byte, 128)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	}
	conn.Close()
	fmt.Println(conn.RemoteAddr().String() + " leave !")
}
