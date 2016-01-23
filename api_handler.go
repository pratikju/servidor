package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	baseResp := BaseResponse{
		CreateRepositoryUrl: fmt.Sprintf(GetProtocol(false) + r.Host + GetRepoCreateUrl()),
		UserRepositoriesUrl: fmt.Sprintf(GetProtocol(false) + r.Host + GetReposUrl()),
		UserRepositoryUrl:   fmt.Sprintf(GetProtocol(false) + r.Host + GetRepoUrl()),
		BranchesUrl:         fmt.Sprintf(GetProtocol(false) + r.Host + GetBranchesUrl()),
		BranchUrl:           fmt.Sprintf(GetProtocol(false) + r.Host + GetBranchUrl()),
	}

	WriteIndentedJson(w, baseResp, "", "  ")
}

func repoCreateHandler(w http.ResponseWriter, r *http.Request) {

}

func repoIndexHandler(w http.ResponseWriter, r *http.Request) {

}

func repoShowHandler(w http.ResponseWriter, r *http.Request) {

}

func branchIndexHandler(w http.ResponseWriter, r *http.Request) {

}

func branchShowHandler(w http.ResponseWriter, r *http.Request) {

}
