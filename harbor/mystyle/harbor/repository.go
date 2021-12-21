package harbor

import (
	"encoding/json"
	"fmt"
	"go-examples/harbor/mystyle/harbor/model"
)

func (c *harborClient) ListRepositories(projectName string, params *model.ListRepositoriesParams) ([]*model.Repository, error) {
	if params == nil {
		params = model.NewListRepositoriesParams()
	}

	url := fmt.Sprintf(urlRepositories, projectName) + params.ToParamString()
	data, err := c.getJson(url, true)
	if err != nil {
		return nil, err
	}

	var repositories []*model.Repository
	if err := json.Unmarshal(data, &repositories); err != nil {
		return nil, err
	}
	return repositories, nil
}
