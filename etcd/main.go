package main

import (
	"context"
	"go-examples/common/config"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
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

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    cfg.Etcd.Username,
		Password:    cfg.Etcd.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), "/test/foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}
