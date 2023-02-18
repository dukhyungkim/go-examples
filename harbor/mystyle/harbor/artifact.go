package harbor

import (
	"encoding/json"
	"fmt"
	"harbor/mystyle/harbor/model"
)

func (c *harborClient) ListArtifacts(projectName string, repositoryName string, params *model.ListArtifactsParams) ([]*model.Artifact, error) {
	if params == nil {
		params = model.NewListArtifactsParams()
	}

	url := fmt.Sprintf(urlArtifacts, projectName, repositoryName) + params.ToParamString()
	data, err := c.getJson(url, true)
	if err != nil {
		return nil, err
	}

	var artifacts []*model.Artifact
	if err := json.Unmarshal(data, &artifacts); err != nil {
		return nil, err
	}
	return artifacts, nil
}
