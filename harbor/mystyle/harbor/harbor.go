package harbor

import (
	"go-examples/harbor/mystyle/harbor/model"
)

type HarborClient interface {
	ListProjects(params *model.ListProjectParams) ([]*model.Project, error)
	Ping() (string, error)
	ListRepositories(projectName string) ([]*model.Repository, error)
	ListArtifacts(projectName string, repositoryName string) ([]*model.Artifact, error)
}
