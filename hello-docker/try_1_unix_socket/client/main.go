package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	check_err := func(err error) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	conn, err := net.Dial("unix", "../test_socket")
	check_err(err)

	in := bufio.NewReader(os.Stdin)
	for true {
		str, _, err := in.ReadLine()
		if err != nil {
			fmt.Println("get msg err-->", err)
			continue
		}
		fmt.Println("sending-->", string(str))
		conn.Write(append(str, '\n'))
		if strings.EqualFold(string(str), "EOF") {
			break
		}
	}

}
