package main

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2"
	"log"
)

func main() {
	const (
		apiURL   = "https://domain/api"
		username = "user"
		password = "pass"
	)
	harborClient, err := apiv2.NewRESTClientForHost(apiURL, username, password, nil)
	if err != nil {
		panic(err)
	}

	projects, err := harborClient.ListProjects(context.Background(), "")
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		log.Println(project.Name)
	}
}
