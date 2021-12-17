package harbor

import (
	"encoding/json"
	"go-examples/harbor/mystyle/harbor/model"
)

func (c *harborClient) ListProjects() ([]*model.Project, error) {
	const projectsURI = urlProjects + "?page=1&page_size=10&with_detail=true"
	data, err := c.getJson(projectsURI, true)
	if err != nil {
		return nil, err
	}

	var projects []*model.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
