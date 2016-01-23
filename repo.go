package main

type Repository struct {
	Name        string   `json:"name"`
	CloneUrl    string   `json:"clone_url"`
	Owner       string   `json:"owner"`
	BranchesUrl []Branch `json:"branches_url"`
}
