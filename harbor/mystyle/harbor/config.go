package harbor

const (
	urlProjects     = "/projects"
	urlPing         = "/ping"
	urlRepositories = "/projects/%s/repositories"
	urlArtifacts    = "/projects/%s/repositories/%s/artifacts"
)

type HarborConfig struct {
	URL      string
	Username string
	Password string
}
