package harbor

const (
	urlProjects     = "/projects"
	urlPing         = "/ping"
	urlRepositories = "/projects/%s/repositories"
)

type HarborConfig struct {
	URL      string
	Username string
	Password string
}
