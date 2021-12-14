package model

type Project struct {
	CreationTime string `json:"creation_time"`
	Metadata     struct {
		Severity             string `json:"severity"`
		PreventVul           string `json:"prevent_vul"`
		RetentionID          string `json:"retention_id"`
		ReuseSysCveAllowlist string `json:"reuse_sys_cve_allowlist"`
		Public               string `json:"public"`
		AutoScan             string `json:"auto_scan"`
		EnableContentTrust   string `json:"enable_content_trust"`
	} `json:"metadata"`
	UpdateTime         string      `json:"update_time"`
	OwnerName          string      `json:"owner_name"`
	ProjectID          int         `json:"project_id"`
	OwnerID            int         `json:"owner_id"`
	Name               string      `json:"name"`
	CurrentUserRoleIDs interface{} `json:"current_user_role_ids"`
	CveAllowlist       struct {
		CreationTime string        `json:"creation_time"`
		UpdateTime   string        `json:"update_time"`
		ProjectID    int           `json:"project_id"`
		ID           int           `json:"id"`
		Items        []interface{} `json:"items"`
	} `json:"cve_allowlist"`
	RepoCount  int `json:"repo_count"`
	ChartCount int `json:"chart_count"`
}
