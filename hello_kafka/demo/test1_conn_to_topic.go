package demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var hosts []string = []string{"0.0.0.0:9092", "0.0.0.0:9093", "0.0.0.0:9094", "0.0.0.0:9095"}

func Conn_to_cluster() {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(hosts, config)
	if err != nil {
		fmt.Println("conn kafka failed-->", err)
		return
	}
	defer client.Close()
	fmt.Println("conn success Brokers-->", len(client.Brokers()))
}

func Create_topic() {
	broker := sarama.NewBroker("0.0.0.0:9093")
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0
	broker.Open(config)
	isConn, err := broker.Connected()
	if err != nil {
		fmt.Println("conn kafka failed-->", err)
		return
	}
	fmt.Println("conn kafka state-->", isConn)
	defer broker.Close()

	topInfo := &sarama.TopicDetail{}
	topInfo.NumPartitions = 2     //两个分区
	topInfo.ReplicationFactor = 3 //三个副本

	topInfos := make(map[string]*sarama.TopicDetail)
	topInfos["fuck_test"] = topInfo

	req := sarama.CreateTopicsRequest{
		TopicDetails: topInfos,
		Timeout:      time.Second * 15,
	}
	resps, err := broker.CreateTopics(&req)
	if err != nil {
		fmt.Println("create topic failed-->", err)
		return
	}
	t := resps.TopicErrors
	for key, val := range t {
		fmt.Println("Key is ", key)
		fmt.Println("Error is ", val.Err.Error())
		fmt.Println("ErrMsg is ", val.ErrMsg)
	}
}
