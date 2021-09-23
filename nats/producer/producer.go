package main

import (
	"github.com/nats-io/nats.go"
	"go-examples/common/config"
	"log"
	"strings"
)

type Producer struct {
	nc *nats.Conn
	js nats.JetStreamContext
}

var Client *Producer

func NewProducer(cfg *config.Nats) error {
	if Client != nil {
		return nil
	}

	nc, err := nats.Connect(strings.Join(cfg.Servers, ","), nats.UserInfo(cfg.Username, cfg.Password))
	if err != nil {
		return err
	}

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	Client = &Producer{nc: nc, js: js}
	return nil
}

func (p *Producer) Close() {
	p.nc.Close()
}

func (p *Producer) CreateStream(streamName string) error {
	streamInfo, err := p.js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	streamSubjects := streamName + ".*"
	if streamInfo == nil {
		log.Printf("create stream %q and subjects %q", streamName, streamSubjects)
		streamConfig := nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		}
		if _, err := p.js.AddStream(&streamConfig); err != nil {
			return err
		}
	}
	return nil
}

func (p *Producer) SendMessage(subject string, message []byte) error {
	pub, err := p.js.Publish(subject, message)
	if err != nil {
		return err
	}

	log.Printf("%+v\n", pub)
	return nil
}
