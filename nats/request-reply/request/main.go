package main

import (
	"go-examples/common/config"
	"log"
	"strings"
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

	req, err := nc.Request(subject, []byte("world"), time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(req.Data))

	nc.Close()
}
