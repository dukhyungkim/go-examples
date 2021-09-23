package main

import (
	"go-examples/common/config"
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

	if err := NewProducer(cfg.Nats); err != nil {
		log.Fatalln(err)
	}
	defer Client.Close()

	if err := Client.CreateStream(cfg.Nats.Subject); err != nil {
		log.Fatalln(err)
	}

	message := "hello"
	if err := Client.SendMessage(cfg.Nats.Subject, []byte(message)); err != nil {
		log.Fatalln(err)
	}
}
