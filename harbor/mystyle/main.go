package main

import (
	"fmt"
	"go-examples/common/config"
	"go-examples/harbor/mystyle/harbor"
	"go-examples/harbor/mystyle/harbor/model"
	"strings"
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

	fmt.Println("--- Ping ---")
	fmt.Println(pong)
	fmt.Println()

	listProjectsParams := model.NewListProjectsParams()
	projects, err := hc.ListProjects(listProjectsParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Projects ---")
	for _, project := range projects {
		fmt.Printf("%+v\n", *project)
	}
	fmt.Println()

	listRepositoriesParams := model.NewListRepositoriesParams()
	repositories, err := hc.ListRepositories(projects[0].Name, listRepositoriesParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Repositories ---")
	for _, repository := range repositories {
		fmt.Printf("%+v\n", *repository)
	}
	fmt.Println()

	artifacts, err := hc.ListArtifacts(projects[0].Name, strings.Split(repositories[0].Name, "/")[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Artifacts ---")
	for _, artifact := range artifacts {
		fmt.Printf("%+v\n", *artifact)
	}
	fmt.Println()
}
