package main

import (
	"fmt"
	"strings"
)

type Repository struct {
	Name        string `json:"name"`
	CloneURL    string `json:"clone_url"`
	Owner       string `json:"owner"`
	BranchesURL string `json:"branches_url"`
}

func GetRepository(h, u, r string) Repository {
	var repo Repository
	rawRepoName := strings.Split(r, ".git")[0]
	repo = Repository{Name: rawRepoName,
		CloneURL:    FormCloneURL(h, u, r),
		Owner:       u,
		BranchesURL: fmt.Sprintf("%s/api/%s/repos/%s/branches{/branch-name}", GetProtocol(config.SSLEnabled)+h, u, rawRepoName),
	}
	return repo
}
