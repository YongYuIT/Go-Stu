package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	check_err := func(err error) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	socket, err := net.ResolveUnixAddr("unix", "./test_socket")
	check_err(err)

	listener, err := net.ListenUnix("unix", socket)
	check_err(err)

	for {
		fmt.Println("wait to accept...")
		conn, err := listener.AcceptUnix()
		clientAddr := conn.RemoteAddr().String()

		fmt.Println("conn with -->", clientAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(conn *net.UnixConn) {
			defer func() {
				fmt.Println(" Disconnected : ", clientAddr)
				conn.Close()
			}()
			fmt.Println("start new gor for : ", conn.RemoteAddr().String())
			reader := bufio.NewReader(conn)
			for true {
				message, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("get err when read-->", err)
					if err == io.EOF {
						break
					} else {
						continue
					}
				}
				fmt.Println("read msg from client-->", message)
			}
		}(conn)
	}
}
