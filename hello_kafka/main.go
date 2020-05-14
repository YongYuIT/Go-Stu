package main

import (
	"flag"
	"fmt"
	"hello_kafka/demo"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("env ok")

	var job_name string
	var proc_name string
	flag.StringVar(&job_name, "j", "hello", "任务名，默认hello")
	flag.StringVar(&proc_name, "p", strconv.FormatInt(time.Now().UnixNano(), 10), "进程标识")
	flag.Parse()

	if strings.EqualFold(job_name, "hello") {
		fmt.Println("hello word")
	}

	topicName := "fuck_test20200513006"

	if strings.EqualFold(job_name, "conn_to_kafka") {
		demo.Conn_to_cluster()
	} else if strings.EqualFold(job_name, "create_topic") {
		demo.Create_topic(topicName)
	} else if strings.EqualFold(job_name, "send_message") {
		demo.Send_message(topicName)
	} else if strings.EqualFold(job_name, "recv_message") {
		//demo.GetMessageFromKafka(topicName)
		demo.GetMessageFromKafkaWithOff(topicName)
	} else if strings.EqualFold(job_name, "recv_message_grp") {
		demo.GetMessageFromKafkaGroup(topicName, proc_name)
	}

}
