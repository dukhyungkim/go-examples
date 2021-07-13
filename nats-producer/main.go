package main

import (
	"log"
	"nats-producer/config"
	"nats-producer/producer"
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

	if err := producer.NewProducer(cfg.Producer); err != nil {
		log.Fatalln(err)
	}
	defer producer.Client.Close()

	if err := producer.Client.CreateStream(cfg.Producer.Subject); err != nil {
		log.Fatalln(err)
	}

	message := "hello"
	if err := producer.Client.SendMessage(cfg.Producer.Subject, []byte(message)); err != nil {
		log.Fatalln(err)
	}
}
