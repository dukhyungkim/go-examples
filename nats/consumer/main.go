package main

import (
	"context"
	"go-examples/common/config"
	"log"
	"os"
	"os/signal"
	"sync"
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

	if err := NewConsumer(cfg.Nats); err != nil {
		log.Fatalf("Cannot init consumer: %v\n", err)
	}
	defer Client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Client.ListenMessage(ctx)
	}()
	log.Println("Consumer up and running!...")

	waitSignal(ctx, cancel)
	wg.Wait()
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
