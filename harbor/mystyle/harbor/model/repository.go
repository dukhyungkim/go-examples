package model

type Repository struct {
	CreationTime  string `json:"creation_time"`
	UpdateTime    string `json:"update_time"`
	ProjectID     int    `json:"project_id"`
	ArtifactCount int    `json:"artifact_count"`
	Name          string `json:"name"`
	ID            int    `json:"id"`
	PullCount     int    `json:"pull_count"`
}
