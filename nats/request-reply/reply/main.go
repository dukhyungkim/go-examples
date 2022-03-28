package main

import (
	"go-examples/common/config"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		defer wg.Done()
		log.Printf("server got message: %s\n", string(msg.Data))
		message := "Hello " + string(msg.Data)
		if err := msg.Respond([]byte(message)); err != nil {
			log.Fatalln(err)
		}
		log.Printf("server sent message: %s\n", message)
	})
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()
	time.Sleep(time.Second)

	if err := sub.Drain(); err != nil {
		log.Fatalln(err)
	}

	if err := nc.Drain(); err != nil {
		log.Fatalln(err)
	}
}
