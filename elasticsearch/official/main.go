package main

import (
	"encoding/json"
	"go-examples/common/config"
	"log"
	"strings"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
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

	client, err := elasticsearch8.NewClient(elasticsearch8.Config{
		Addresses: cfg.Elasticsearch.Addresses,
		Username:  cfg.Elasticsearch.Username,
		Password:  cfg.Elasticsearch.Password,
	})
	if err != nil {
		log.Panicln(err)
	}

	resp, err := client.Info()
	if err != nil {
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.IsError() {
		log.Fatalf("Error: %s", resp.String())
	}

	var r map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch8.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))
}
