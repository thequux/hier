package common

import (
	"fmt"
	git "github.com/libgit2/git2go"
)

const HierBranchPrefix string = "refs/hier"

type AppData struct {
	Repo *git.Repository
}

func OpenRepo(path string) (*AppData, error) {
	var err error
	var repo *git.Repository
	
	if path == "" {
		path, err = git.Discover(".", false, nil)
		if err != nil {
			return nil, err
		}
	}
	repo, err = git.OpenRepository(path)
	if err != nil {
		return nil, err
	}

	// Check for the safety branch to prevent beta code from
	// running on its own repo
	var verboten bool
	var config *git.Config
	config, err = repo.Config()
	if err != nil {
		return nil, err
	}
	verboten, err = config.LookupBool("hier.verboten")
	if err == nil && verboten {
		return nil, fmt.Errorf("Refusing to open repo at %s because hier.verboten is set", path)
	}
	
	return &AppData{Repo: repo}, nil
}
