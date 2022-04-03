package main

import (
	"context"
	"go-examples/common/config"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Panicln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalf("Cannot access config: %v\n", err)
	}

	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.Elasticsearch.Addresses...),
		elastic.SetBasicAuth(cfg.Elasticsearch.Username, cfg.Elasticsearch.Password),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Panicln(err)
	}

	info, code, err := client.Ping(cfg.Elasticsearch.Addresses[0]).Do(ctx)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("ClusterName: %s\n", info.ClusterName)
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
