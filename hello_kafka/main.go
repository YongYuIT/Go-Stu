package main

import (
	"fmt"
	"hello_kafka/demo"
)

func main() {
	fmt.Println("env ok")
	//demo.Conn_to_cluster()
	//demo.Create_topic()
	demo.Send_message()
}
