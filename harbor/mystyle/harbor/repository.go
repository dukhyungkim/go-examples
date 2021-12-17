package harbor

import (
	"encoding/json"
	"fmt"
	"go-examples/harbor/mystyle/harbor/model"
)

func (c *harborClient) ListRepositories(projectName string) ([]*model.Repository, error) {
	data, err := c.getJson(fmt.Sprintf(urlRepositories, projectName), true)
	if err != nil {
		return nil, err
	}

	var repositories []*model.Repository
	if err := json.Unmarshal(data, &repositories); err != nil {
		return nil, err
	}
	return repositories, nil
}
