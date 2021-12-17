package harbor

import (
	"go-examples/harbor/mystyle/harbor/model"
)

type HarborClient interface {
	ListProjects() ([]*model.Project, error)
	Ping() (string, error)
}
