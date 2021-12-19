package harbor

import (
	"encoding/json"
	"fmt"
	"go-examples/harbor/mystyle/harbor/model"
)

func (c *harborClient) ListArtifacts(projectName string, repositoryName string) ([]*model.Artifact, error) {
	data, err := c.getJson(fmt.Sprintf(urlArtifacts, projectName, repositoryName), true)
	if err != nil {
		return nil, err
	}

	var artifacts []*model.Artifact
	if err := json.Unmarshal(data, &artifacts); err != nil {
		return nil, err
	}
	return artifacts, nil
}
