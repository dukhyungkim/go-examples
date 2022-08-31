package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	HOST  = "192.168.0.219"
	TOPIC = "myTopic"
)

func main() {
	producer := kafka.Writer{
		Addr:  kafka.TCP(HOST),
		Topic: TOPIC,
	}
	defer func(producer *kafka.Writer) {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}(&producer)

	ctx := context.Background()

	var input string
	var err error
	for {
		if _, err = fmt.Scan(&input); err != nil {
			panic(err)
		}

		if err = producer.WriteMessages(ctx,
			kafka.Message{Value: []byte(input)},
		); err != nil {
			panic(err)
		}
	}
}
