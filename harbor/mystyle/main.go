package main

import (
	"go-examples/common/config"
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

	hc := NewHarborClient(&HarborConfig{
		URL:      cfg.Harbor.APIHost,
		Username: cfg.Harbor.Username,
		Password: cfg.Harbor.Password,
	})

	err = hc.ListProjects()
	if err != nil {
		log.Println(err)
	}
}
