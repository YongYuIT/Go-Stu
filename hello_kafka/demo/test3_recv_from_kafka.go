package demo

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func GetMessageFromKafka(topicName string) {
	consumer, err := sarama.NewConsumer([]string{"kafka1.thinking.com:9092", "kafka2.thinking.com:9092", "kafka3.thinking.com:9092", "kafka4.thinking.com:9092"}, nil)
	if err != nil {
		fmt.Println("conn kafka failed-->", err)
		return
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(topicName)
	if err != nil {
		fmt.Println("cannot get parts-->", err)
		return
	}

	fmt.Println("find parts-->", len(partitions))

	for i := 0; i < len(partitions); i++ {
		partitionConsumer, err := consumer.ConsumePartition(topicName, partitions[i], sarama.OffsetNewest)
		if err != nil {
			fmt.Println("get part consu failed-->", err)
			continue
		}
		//每个分区一个协程去读取消息
		fmt.Println("start read at-->", i)
		go_index := i
		go func() {
			for msg := range partitionConsumer.Messages() {
				//实验表明，偏移量在特定主题特定分区内是唯一的，单调递增的
				fmt.Println("recv-->", go_index, "-->", string(msg.Key), ":", string(msg.Value), "-->", msg.Offset)
			}
		}()
	}

	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
}
