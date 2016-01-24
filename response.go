package main

type BaseResponse struct {
	CreateRepositoryURL string `json:"create_repo_url"`
	UserRepositoriesURL string `json:"user_repositories_url"`
	UserRepositoryURL   string `json:"user_repository_url"`
	BranchesURL         string `json:"branches_url"`
}

type CreateResponse struct {
	ResponseMessage string `json:"response_message"`
	CloneURL        string `json:"clone_url"`
}
