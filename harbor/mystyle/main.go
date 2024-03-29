package main

import (
	"fmt"
	"harbor/config"
	"harbor/mystyle/harbor"
	"harbor/mystyle/harbor/model"
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

	projectName := projects[0].Name

	listRepositoriesParams := model.NewListRepositoriesParams()
	repositories, err := hc.ListRepositories(projectName, listRepositoriesParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Repositories ---")
	for _, repository := range repositories {
		fmt.Printf("%+v\n", *repository)
	}
	fmt.Println()

	repositoryName := strings.Split(repositories[0].Name, "/")[1]

	listArtifactsParams := model.NewListArtifactsParams()
	artifacts, err := hc.ListArtifacts(projectName, repositoryName, listArtifactsParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Artifacts ---")
	for _, artifact := range artifacts {
		fmt.Printf("%+v\n", *artifact)
	}
	fmt.Println()
}
