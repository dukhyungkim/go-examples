package harbor

const (
	urlProjects = "/projects"
	urlPing     = "/ping"
)

type HarborConfig struct {
	URL      string
	Username string
	Password string
}
