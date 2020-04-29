package demo

import (
	"fmt"
	"github.com/Shopify/sarama"
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
