package main

import (
	"github.com/nats-io/nats.go"
	"go-examples/common/config"
	"log"
	"strings"
)

const (
	subject = "request-reply"
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

	nc, err := nats.Connect(strings.Join(cfg.Nats.Servers, ","), nats.UserInfo(cfg.Nats.Username, cfg.Nats.Password))
	if err != nil {
		log.Fatalln(err)
	}

	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		message := "Hello " + string(msg.Data)
		if err := msg.Respond([]byte(message)); err != nil {
			log.Fatalln(err)
		}
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err := sub.Unsubscribe(); err != nil {
		log.Fatalln(err)
	}

	nc.Close()
}
