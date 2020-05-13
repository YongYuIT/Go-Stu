package main

import (
	"flag"
	"fmt"
	"hello_kafka/demo"
	"strings"
)

func main() {
	fmt.Println("env ok")

	var job_name string
	flag.StringVar(&job_name, "j", "hello", "任务名，默认hello")
	flag.Parse()

	if strings.EqualFold(job_name, "hello") {
		fmt.Println("hello word")
	}

	topicName := "fuck_test20200509004"

	if strings.EqualFold(job_name, "conn_to_kafka") {
		demo.Conn_to_cluster()
	} else if strings.EqualFold(job_name, "create_topic") {
		demo.Create_topic(topicName)
	} else if strings.EqualFold(job_name, "send_message") {
		demo.Send_message(topicName)
	} else if strings.EqualFold(job_name, "recv_message") {
		//demo.GetMessageFromKafka(topicName)
		demo.GetMessageFromKafkaWithOff(topicName)
	}

}
