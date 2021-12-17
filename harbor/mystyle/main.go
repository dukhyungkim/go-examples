package main

import (
	"fmt"
	"go-examples/common/config"
	"go-examples/harbor/mystyle/harbor"
	"log"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		panic(err)
	}

	hc := harbor.NewHarborClient(&harbor.HarborConfig{
		URL:      cfg.Harbor.APIHost,
		Username: cfg.Harbor.Username,
		Password: cfg.Harbor.Password,
	})

	pong, err := hc.Ping()
	if err != nil {
		panic(err)
	}
	log.Println(pong)
	fmt.Println()

	projects, err := hc.ListProjects()
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		log.Printf("%+v\n", *project)
	}
	fmt.Println()

	repositories, err := hc.ListRepositories(projects[0].Name)
	if err != nil {
		panic(err)
	}

	for _, repository := range repositories {
		log.Printf("%+v\n", *repository)
	}
	fmt.Println()
}
