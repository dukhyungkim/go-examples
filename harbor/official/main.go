package main

import (
	"context"
	"go-examples/common/config"
	"log"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
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

	c := &harbor.ClientSetConfig{
		URL:      "https://harbor.brique.kr",
		Password: cfg.Harbor.Username,
		Username: cfg.Harbor.Password,
	}

	cs, err := harbor.NewClientSet(c)
	if err != nil {
		log.Fatalln(err)
	}

	getHealth, err := cs.V2().Health.GetHealth(context.Background(), health.NewGetHealthParams())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(getHealth.GetPayload().Status)

	params := repository.NewListRepositoriesParams()
	params.SetProjectName("qps")
	repositories, err := cs.V2().Repository.ListRepositories(context.Background(), params)
	if err != nil {
		log.Fatalln(err)
	}

	for _, repo := range repositories.GetPayload() {
		log.Printf("id: %d, name: %s, artifactCount: %d, pullCount: %d, createdAt: %s, updatedAt: %s, desc: %s\n",
			repo.ID, repo.Name, repo.ArtifactCount, repo.PullCount, repo.CreationTime, repo.UpdateTime, repo.Description)
	}
}
