package main

import (
	"go-examples/common/config"
	"go-examples/nats-producer/producer"
	"log"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalf("Cannot access config: %v\n", err)
	}

	if err := producer.NewProducer(cfg.Nats); err != nil {
		log.Fatalln(err)
	}
	defer producer.Client.Close()

	if err := producer.Client.CreateStream(cfg.Nats.Subject); err != nil {
		log.Fatalln(err)
	}

	message := "hello"
	if err := producer.Client.SendMessage(cfg.Nats.Subject, []byte(message)); err != nil {
		log.Fatalln(err)
	}
}
