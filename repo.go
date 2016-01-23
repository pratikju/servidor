package main

import (
	"fmt"
	"strings"
)

type Repository struct {
	Name        string `json:"name"`
	CloneUrl    string `json:"clone_url"`
	Owner       string `json:"owner"`
	BranchesUrl string `json:"branches_url"`
}

func GetRepository(h, u, r string) Repository {
	var repo Repository
	rawRepoName := strings.Split(r, ".git")[0]
	repo = Repository{Name: rawRepoName,
		CloneUrl:    FormCloneURL(h, u, r),
		Owner:       u,
		BranchesUrl: fmt.Sprintf("%s/api/%s/repos/%s/branches{/branch-name}", GetProtocol(false)+h, u, rawRepoName),
	}
	return repo
}
