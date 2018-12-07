package github

import (
	"time"

	"gitlab.com/rosenpin/git-project-showcaser/api/services"
	"gitlab.com/rosenpin/git-project-showcaser/api/utils"
	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// Github is a git service that projects can be fetched from
type Github struct {
	*GithubFetcher
}

// NewGithub creates a new Github object and uses the timeout as the request timeout for API requests
func NewGithub(timeout time.Duration) services.Service {
	return &Github{NewGithubFetcher(utils.NewHTTPJsonFetcher(timeout))}
}

func (github *Github) GetProjectsFor(username string) ([]*models.Project, error) {
	return github.FetchProjects(username)
}