package demo

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strconv"
	"strings"
	"time"
)

var kafka_address = []string{"kafka1.thinking.com:9092", "kafka2.thinking.com:9092", "kafka3.thinking.com:9092", "kafka4.thinking.com:9092"}
var zk_address = []string{"test_1.thinking.com:2181", "test_2.thinking.com:2181", "test_3.thinking.com:2181"}

func GetMessageFromKafka(topicName string) {
	consumer, err := sarama.NewConsumer(kafka_address, nil)
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
	consumer, err := sarama.NewConsumer(kafka_address, nil)
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

	conn, _, err := zk.Connect(zk_address, 10*time.Second)
	if err != nil {
		fmt.Println("zk conn error-->", err)
		return
	}
	defer conn.Close()
	fmt.Println("zk conn success!")

	path := "/" + topicName + "_off"
	exist, state, err := conn.Exists(path)
	if err != nil {
		fmt.Println("read path state err-->", err, state.Version)
		return
	}
	if !exist {
		conn.Create(path, []byte("init_val"), 0, zk.WorldACL(zk.PermAll))
	}

	for pindex := 0; pindex < len(partitions); pindex++ {
		sub_path := path + "/part_" + strconv.Itoa(pindex)
		exist, state, err := conn.Exists(sub_path)
		if err != nil {
			fmt.Println("read path state err-->", err, state.Version)
			return
		}
		if !exist {
			conn.Create(sub_path, []byte("init_val"), 0, zk.WorldACL(zk.PermAll))
		}

		off_data, state, err := conn.Get(sub_path)
		if err != nil {
			fmt.Println("read zk data err-->", err, state.Version, off_data)
			return
		}

		var kafka_start_off int64 = 0
		if strings.EqualFold(string(off_data), "init_val") {
			kafka_start_off = sarama.OffsetNewest
		} else {
			kafka_start_off, err = strconv.ParseInt(string(off_data), 10, 64)
			kafka_start_off += 1 //加一很重要，避免重复消费
			if err != nil {
				fmt.Println("read zk off data err-->", err, state.Version, off_data)
			}
		}

		fmt.Println("get read start point-->", pindex, "-->", kafka_start_off)
		partitionConsumer, err := consumer.ConsumePartition(topicName, partitions[pindex], kafka_start_off)
		if err != nil {
			fmt.Println("get part consu failed-->", err)
			continue
		}
		//每个分区一个协程去读取消息
		fmt.Println("start read at part(outside)-->", pindex)
		go_index := pindex
		go func() {
			fmt.Println("start read at part-->", go_index, partitionConsumer == nil)
			for msg := range partitionConsumer.Messages() {
				zk_data, state, err := conn.Get(sub_path)
				fmt.Println("zk before-->", string(zk_data))
				state, err = conn.Set(sub_path, []byte(strconv.FormatInt(msg.Offset, 10)), state.Version)
				if err != nil {
					fmt.Println("write zk err --> ", err)
				}
				fmt.Println("recv success-->", go_index, "-->", msg.Offset, "-->", string(msg.Key), "-->", string(msg.Value))
				fmt.Println("write zk success-->", state.Version, "-->", msg.Offset)
			}
		}()
	}

	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
}

func GetMessageFromKafkaGroup(topicName string, p_name string) {
	conf := sarama.NewConfig()
	conf.Version = sarama.V2_4_0_0
	//决定分区分配策略，即组内哪个消费者得到哪个分区
	conf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	//默认的组消费偏移
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	//消费组名称
	group_name := "fuck_test_003_grp"
	group, err := sarama.NewConsumerGroup(kafka_address, group_name, conf)
	if err != nil {
		fmt.Println("create err-->", err, "-->", p_name, "-->", group == nil)
	} else {
		fmt.Println("create success-->", p_name)
	}
	ctx, cancel := context.WithCancel(context.Background())

	//此方法为阻塞方法，在重平衡（组员发生变化）时会退出。需要优化，避免重平衡后直接停止了消费。
	err = group.Consume(ctx, []string{topicName}, myHandler{p_name, topicName})
	if err != nil {
		fmt.Println("cus init err-->", err)
	} else {
		fmt.Println("cus init success")
	}

	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
	cancel()
}

type myHandler struct {
	Pid   string
	Topic string
}

func (this myHandler) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (this myHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil

}

func (this myHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, cla sarama.ConsumerGroupClaim) error {
	for msg := range cla.Messages() {
		fmt.Println("recv success-->", this.Pid, "-->", msg.Partition, "-->", msg.Offset, "-->", string(msg.Key), "-->", string(msg.Value))
		// 更新位移
		sess.MarkMessage(msg, "")
	}
	return nil
}
