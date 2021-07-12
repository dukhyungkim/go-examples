package main

import (
	"context"
	"log"
	"nats-consumer/config"
	"nats-consumer/consumer"
	"os"
	"os/signal"
	"syscall"
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

	if err := consumer.NewConsumer(cfg.Consumer); err != nil {
		log.Fatalf("Cannot init consumer: %v\n", err)
	}
	defer consumer.Client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go consumer.Client.ListenMessage(ctx)
	log.Println("Consumer up and running!...")

	waitSignal(ctx, cancel)
}

func waitSignal(ctx context.Context, cancel context.CancelFunc) {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}

	cancel()
}
