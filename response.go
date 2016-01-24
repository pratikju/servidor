package main

type BaseResponse struct {
	CreateRepositoryUrl string `json:"create_repo_url"`
	UserRepositoriesUrl string `json:"user_repositories_url"`
	UserRepositoryUrl   string `json:"user_repository_url"`
	BranchesUrl         string `json:"branches_url"`
}

type CreateResponse struct {
	ResponseMessage string `json:"response_message"`
	CloneUrl        string `json:"clone_url"`
}
