package main

import (
	"encoding/json"
	"fmt"
	"github.com/libgit2/git2go"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Payload struct {
	Username string
	RepoName string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	baseResp := BaseResponse{
		CreateRepositoryURL: fmt.Sprintf(GetProtocol(config.SSLEnabled) + r.Host + GetRepoCreateURL()),
		UserRepositoriesURL: fmt.Sprintf(GetProtocol(config.SSLEnabled) + r.Host + GetReposURL()),
		UserRepositoryURL:   fmt.Sprintf(GetProtocol(config.SSLEnabled) + r.Host + GetRepoURL()),
		BranchesURL:         fmt.Sprintf(GetProtocol(config.SSLEnabled) + r.Host + GetBranchesURL()+"{/branch-name}"),
	}

	WriteIndentedJSON(w, baseResp, "", "  ")
}

func repoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var resp CreateResponse
	resp.ResponseMessage = "Unknown error. Follow README"
	resp.CloneURL = ""

	wd, _ := os.Getwd()

	defer func() {
		WriteIndentedJSON(w, resp, "", "  ")
		if err := os.Chdir(wd); err != nil {
			log.Println(err)
		}
	}()

	var payload Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
		return
	}
	if payload.Username == "" || payload.RepoName == "" {
		log.Println("Empty username or reponame")
		return
	}

	usrPath := UserPath(payload.Username)
	bareRepo := FormatRepoName(payload.RepoName)
	url := FormCloneURL(r.Host, payload.Username, bareRepo)

	if ok := IsExistingRepository(RepoPath(payload.Username, payload.RepoName)); ok {
		resp.ResponseMessage = fmt.Sprintf("repository already exists for %s", payload.Username)
		resp.CloneURL = url
		return
	}

	if err := os.MkdirAll(usrPath, 0775); err != nil {
		resp.ResponseMessage = "error while creating user"
		return
	}

	if err := os.Chdir(usrPath); err != nil {
		resp.ResponseMessage = "error while creating new repository"
		return
	}

	cmd := exec.Command(config.GitPath, "init", "--bare", bareRepo)

	if err := cmd.Start(); err == nil {
		resp.CloneURL = url
		resp.ResponseMessage = "repository created successfully"
	} else {
		resp.ResponseMessage = "error while creating new repository"
		return
	}
	if err := cmd.Wait(); err != nil {
		log.Println("Error while waiting:", err)
		return
	}
}

func repoIndexHandler(w http.ResponseWriter, r *http.Request) {
	userName, _, _ := GetParamValues(r)
	var errJSON Error
	list, ok := FindAllDir(UserPath(userName))
	if !ok {
		errJSON = Error{Message: "repository not found"}
		WriteIndentedJSON(w, errJSON, "", "  ")
		return
	}
	var repo Repository
	var repos []Repository

	for i := 0; i < len(list); i++ {
		repo = GetRepository(r.Host, userName, list[i].Name())
		repos = append(repos, repo)
	}
	WriteIndentedJSON(w, repos, "", "  ")
}

func repoShowHandler(w http.ResponseWriter, r *http.Request) {
	var errJSON Error
	userName, repoName, _ := GetParamValues(r)
	if ok := IsExistingRepository(RepoPath(userName, repoName)); !ok {
		errJSON = Error{Message: "repository not found"}
		WriteIndentedJSON(w, errJSON, "", "  ")
		return
	}
	repo := GetRepository(r.Host, userName, FormatRepoName(repoName))
	WriteIndentedJSON(w, repo, "", "  ")
}

func branchIndexHandler(w http.ResponseWriter, r *http.Request) {
	var errJSON Error
	userName, repoName, _ := GetParamValues(r)
	if ok := IsExistingRepository(RepoPath(userName, repoName)); !ok {
		errJSON = Error{Message: "repository not found"}
		WriteIndentedJSON(w, errJSON, "", "  ")
		return
	}
	re, _ := git.OpenRepository(RepoPath(userName, repoName))
	branches, _ := GetBranches(re)
	WriteIndentedJSON(w, branches, "", "  ")
}

func branchShowHandler(w http.ResponseWriter, r *http.Request) {
	var errJSON Error
	userName, repoName, branchName := GetParamValues(r)
	if ok := IsExistingRepository(RepoPath(userName, repoName)); !ok {
		errJSON = Error{Message: "repository not found"}
		WriteIndentedJSON(w, errJSON, "", "  ")
		return
	}

	re, _ := git.OpenRepository(RepoPath(userName, repoName))
	branch, ok := GetBranchByName(branchName, re)
	if !ok {
		errJSON = Error{Message: "branch not found"}
		WriteIndentedJSON(w, errJSON, "", "  ")
		return
	}

	WriteIndentedJSON(w, branch, "", "  ")
}
