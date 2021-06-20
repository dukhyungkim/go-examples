package main

import (
	"context"
	"log"
	"nats-consumer/consumer"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	c, err := consumer.NewConsumer()
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if ctx.Err() != nil {
				return
			}
			if err := c.ClaimMessage(); err != nil {
				log.Println(err)
				return
			}
		}
	}()

	log.Println("Consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	c.Unsubscribe()
}
