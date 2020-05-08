package demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
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

	off := int64(-1)
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = topname
		send_key := strconv.FormatInt(time.Now().UnixNano(), 10) + "-key"
		fmt.Println("send key-->", send_key)
		msg.Key = sarama.StringEncoder(send_key)
		send_value := "this is a message No. is " + strconv.Itoa(rand.Intn(1000000))
		msg.Value = sarama.StringEncoder(send_value)

		part_id, _off, err := client.SendMessage(msg)
		off = _off
		if err != nil {
			fmt.Println("send error-->", err)
			return
		}
		fmt.Println("send success-->", part_id, "-->", off, "-->", send_key)
		time.Sleep(5 * time.Second)
	}
}
