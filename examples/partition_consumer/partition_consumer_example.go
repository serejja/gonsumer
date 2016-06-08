package main

import (
	"fmt"
	"github.com/serejja/gonsumer"
	"github.com/serejja/siesta"
)

func main() {
	config := siesta.NewConnectorConfig()
	config.BrokerList = []string{"localhost:9092"}

	client, err := siesta.NewDefaultConnector(config)
	if err != nil {
		panic(err)
	}

	consumer := gonsumer.NewPartitionConsumer(client, gonsumer.NewConsumerConfig(), "gonsumer", 0, partitionConsumerStrategy)

	consumer.Start()
}

func partitionConsumerStrategy(data *gonsumer.FetchData, consumer *gonsumer.KafkaPartitionConsumer) {
	if data.Error != nil {
		panic(data.Error)
	}

	for _, msg := range data.Messages {
		fmt.Printf("%s from partition %d\n", string(msg.Value), msg.Partition)
	}
}
