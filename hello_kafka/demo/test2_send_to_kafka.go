package demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

func Send_message(topname string) {
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
		msg.Topic = topname
		fmt.Println("send key-->", strconv.Itoa(off)+"key")
		msg.Key = sarama.StringEncoder(strconv.Itoa(off) + "key")
		msg.Value = sarama.StringEncoder(strconv.Itoa(off) + "-->fuck kafka test message")

		part_id, _off, err := client.SendMessage(msg)
		off = _off
		if err != nil {
			fmt.Println("send error-->", err)
			return
		}
		fmt.Println("send success-->", part_id, off)
		time.Sleep(5 * time.Second)
	}
}
