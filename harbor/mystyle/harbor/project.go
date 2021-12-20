package harbor

import (
	"encoding/json"
	"go-examples/harbor/mystyle/harbor/model"
)

func (c *harborClient) ListProjects(params *model.ListProjectParams) ([]*model.Project, error) {
	if params == nil {
		params = model.NewListProjectsParams()
	}

	url := urlProjects + params.ToParamString()
	println(url)
	data, err := c.getJson(url, true)
	if err != nil {
		return nil, err
	}

	var projects []*model.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
