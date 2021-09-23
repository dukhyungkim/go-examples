package main

import (
	"context"
	"github.com/nats-io/nats.go"
	"go-examples/common/config"
	"log"
	"strings"
)

type Consumer struct {
	nc  *nats.Conn
	js  nats.JetStreamContext
	sub *nats.Subscription
}

var Client *Consumer

func NewConsumer(cfg *config.Nats) error {
	if Client != nil {
		return nil
	}

	nc, err := nats.Connect(strings.Join(cfg.Servers, ","), nats.UserInfo(cfg.Username, cfg.Password))
	if err != nil {
		return err
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	sub, err := js.PullSubscribe(cfg.Subject, cfg.Group)
	if err != nil {
		return err
	}

	Client = &Consumer{nc: nc, js: js, sub: sub}
	return nil
}

func (c *Consumer) Close() {
	if err := c.sub.Drain(); err != nil {
		log.Println(err)
	}
	c.nc.Close()
}

func (c *Consumer) ListenMessage(ctx context.Context) {
	for {
		if err := c.claimMessage(); err != nil {
			log.Printf("Error from consumer: %v\n", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func (c *Consumer) claimMessage() error {
	messages, err := c.sub.Fetch(1)
	if err == nats.ErrTimeout {

	} else if err != nil {
		return err
	}

	for _, message := range messages {
		if err := message.Ack(); err != nil {
			return err
		}
		log.Printf("Data: %s\n", string(message.Data))
	}
	return nil
}
