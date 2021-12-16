package main

import (
	"go-examples/common/config"
	"go-examples/harbor/mystyle/harbor"
	"log"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Println(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Println(err)
	}

	hc := harbor.NewHarborClient(&harbor.HarborConfig{
		URL:      cfg.Harbor.APIHost,
		Username: cfg.Harbor.Username,
		Password: cfg.Harbor.Password,
	})

	err = hc.ListProjects()
	if err != nil {
		log.Println(err)
	}
}
