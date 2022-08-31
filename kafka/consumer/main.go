package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	HOST  = "192.168.0.219"
	GROUP = "myGroup"
	TOPIC = "myTopic"
)

func main() {
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{HOST},
		GroupID: GROUP,
		Topic:   TOPIC,
	})
	defer func(consumer *kafka.Reader) {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}(consumer)

	fmt.Println("start consuming...")
	ctx := context.Background()
	var message kafka.Message
	var err error
	for {
		message, err = consumer.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message at topic: %s, partition: %d, offset: %d => %s\n", message.Topic, message.Partition, message.Offset, string(message.Value))
	}
}
