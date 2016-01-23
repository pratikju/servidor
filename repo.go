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

	repo = Repository{Name: strings.Split(r, ".git")[0],
		CloneUrl:    FormCloneURL(h, u, r),
		Owner:       u,
		BranchesUrl: fmt.Sprintf(GetProtocol(false) + r + GetBranchesUrl()),
	}
	return repo
}
