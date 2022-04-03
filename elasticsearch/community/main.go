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
	)
	if err != nil {
		log.Panicln(err)
	}

	infoResponse, err := client.NodesInfo().Do(ctx)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("ClusterName: %s\n", infoResponse.ClusterName)
	log.Printf("Nodes: %#+v\n", infoResponse.Nodes)
}
