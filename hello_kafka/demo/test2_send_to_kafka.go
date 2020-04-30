package demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

func Send_message() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"kafka1.thinking.com:9092", "kafka2.thinking.com:9092", "kafka3.thinking.com:9092", "kafka4.thinking.com:9092"}, config)
	if err != nil {
		fmt.Println("create prod err-->", err)
		return
	}
	defer client.Close()

	off := -1
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "fuck_test"
		msg.Value = sarama.StringEncoder("fuck kafka test message" + strconv.Itoa(off))

		pid, off, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send error-->", err)
			return
		}
		fmt.Println("send success-->", pid, off)
		time.Sleep(2 * time.Second)
	}
}
