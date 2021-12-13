package main

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2"
	"go-examples/common/config"
	"log"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	harborClient, err := apiv2.NewRESTClientForHost(cfg.Harbor.APIHost, cfg.Harbor.Username, cfg.Harbor.Password, nil)
	if err != nil {
		panic(err)
	}

	projects, err := harborClient.ListProjects(context.Background(), "qps")
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		log.Println(project.Name)
	}
}
