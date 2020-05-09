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
		//sarama.OffsetNewest决定了消费者每次都是从头消费消息
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
				fmt.Println("recv success-->", go_index, "-->", msg.Offset, "-->", string(msg.Key), "-->", string(msg.Value))
			}
		}()
	}

	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
}

//上面的GetMessageFromKafka方法没有考虑断点问题。这个方法将消费的标识存到zk里面，后面每次启动从上次的标识开始接收数据
func GetMessageFromKafkaWithOff(topicName string) {

}

func GetMessageFromKafkaGroup(topicName string) {
	conf := sarama.NewConfig()
	//决定分区分配策略，即组内哪个消费者得到哪个分区
	conf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
}
