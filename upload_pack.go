package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func uploadPackHandler(w http.ResponseWriter, r *http.Request) {
	userName, repoName, _ := GetParamValues(r)
	execPath := RepoPath(userName, repoName)

	cmd := exec.Command(config.GitPath, "upload-pack", "--stateless-rpc", execPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("Error with child stdin pipe:", err)
		http.Error(w, "Error with child stdin pipe", http.StatusInternalServerError)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error with child stdout pipe:", err)
		http.Error(w, "Error with child stdout pipe", http.StatusInternalServerError)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error with child stderr pipe:", err)
		http.Error(w, "Error with child stderr pipe", http.StatusInternalServerError)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		http.Error(w, "Error while spawning", http.StatusInternalServerError)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading request body:", err)
		http.Error(w, "Error while reading request body", http.StatusInternalServerError)
		return
	}
	stdin.Write(reqBody)

	content_type := "application/x-git-upload-pack-result"
	SetHeader(w, content_type)

	go io.Copy(w, stdout)
	go io.Copy(w, stderr)

	if err := cmd.Wait(); err != nil {
		log.Println("Error while waiting:", err)
		return
	}
}
