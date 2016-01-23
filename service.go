package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	userName, repoName, _ := GetParamValues(r)
	service := FindService(r)
	execPath := RepoPath(userName, repoName)
	if ok := IsExistingRepository(execPath); !ok {
		log.Println("repository not found")
		http.Error(w, "repository not found", http.StatusNotFound)
		return
	}

	cmd := exec.Command(config.GitPath, service, "--stateless-rpc", "--advertise-refs", execPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error with child stdout pipe:", err)
		http.Error(w, "Error with child stdout pipe:", http.StatusInternalServerError)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error with child stderr pipe:", err)
		http.Error(w, "Error with child stderr pipe:", http.StatusInternalServerError)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("Error while spawning:", err)
		http.Error(w, "Error while spawning", http.StatusInternalServerError)
		return
	}

	content_type := fmt.Sprintf("application/x-git-%s-advertisement", service)
	SetHeader(w, content_type)
	w.Write([]byte(CreateFirstPKTLine(service)))
	go io.Copy(w, stdout)
	go io.Copy(w, stderr)
	if err := cmd.Wait(); err != nil {
		log.Println("Error while waiting:", err)
		return
	}
}
