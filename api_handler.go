package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"os/exec"
)

type Payload struct {
	Username string
	RepoName string
}

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
	var resp CreateResponse
	resp.ResponseMessage = "Unknown error. Follow README"
	resp.CloneUrl = ""

	wd, _ := os.Getwd()

	defer func() {
		WriteIndentedJson(w, resp, "", "  ")
		if err := os.Chdir(wd); err != nil {
			log.Println(err)
		}
	}()

	var payload Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
		return
	} else {
		if payload.Username == "" || payload.RepoName == "" {
			log.Println("Empty username or reponame")
			return
		}
	}

	usrPath := UserPath(payload.Username)
	bareRepo := FormatRepoName(payload.RepoName)
	url := fmt.Sprintf(GetProtocol(false) + r.Host + "/" + payload.Username + "/" + bareRepo)

	if _, err := os.Stat(RepoPath(payload.Username, payload.RepoName)); err == nil {
		resp.ResponseMessage = fmt.Sprintf("repository already exists for %s", payload.Username)
		resp.CloneUrl = url
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
		resp.CloneUrl = url
		resp.ResponseMessage = "Repository created successfully"
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

}

func repoShowHandler(w http.ResponseWriter, r *http.Request) {

}

func branchIndexHandler(w http.ResponseWriter, r *http.Request) {

}

func branchShowHandler(w http.ResponseWriter, r *http.Request) {

}
