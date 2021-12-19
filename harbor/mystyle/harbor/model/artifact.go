package model

type Artifact struct {
	AdditionLinks struct {
		Vulnerabilities struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"vulnerabilities"`
		BuildHistory struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"build_history"`
	} `json:"addition_links"`
	References interface{} `json:"references"`
	ExtraAttrs struct {
		Os      string `json:"os"`
		Author  string `json:"author"`
		Created string `json:"created"`
		Config  struct {
			WorkingDir string   `json:"WorkingDir"`
			Entrypoint []string `json:"Entrypoint"`
			Env        []string `json:"Env"`
		} `json:"config"`
		Architecture string `json:"architecture"`
	} `json:"extra_attrs"`
	Icon              string      `json:"icon"`
	ManifestMediaType string      `json:"manifest_media_type"`
	Type              string      `json:"type"`
	Labels            interface{} `json:"labels"`
	Tags              []struct {
		PullTime     string `json:"pull_time"`
		Immutable    bool   `json:"immutable"`
		Name         string `json:"name"`
		Signed       bool   `json:"signed"`
		RepositoryID int    `json:"repository_id"`
		ID           int    `json:"id"`
		ArtifactID   int    `json:"artifact_id"`
		PushTime     string `json:"push_time"`
	} `json:"tags"`
	PullTime     string `json:"pull_time"`
	Size         int    `json:"size"`
	MediaType    string `json:"media_type"`
	ProjectID    int    `json:"project_id"`
	Digest       string `json:"digest"`
	RepositoryID int    `json:"repository_id"`
	ID           int    `json:"id"`
	PushTime     string `json:"push_time"`
}
