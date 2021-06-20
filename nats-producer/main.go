package main

import (
	"log"
	"nats-producer/producer"
)

func main() {
	p, err := producer.NewProducer()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	const streamName = "TEST"
	if err := p.CreateStream(streamName); err != nil {
		log.Fatalln(err)
	}

	if err := p.SendMessage(); err != nil {
		log.Fatalln(err)
	}
}
