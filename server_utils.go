package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetParamValues(r *http.Request) (string, string, string) {
	vars := mux.Vars(r)
	userName := vars["user-name"]
	repoName := vars["repo-name"]
	branchName := vars["branch-name"]
	return userName, repoName, branchName
}

func FindService(r *http.Request) string {
	s := r.URL.Query().Get("service")
	service := strings.SplitN(s, "-", 2)[1]
	return service
}

func SetHeader(w http.ResponseWriter, content_type string) {
	w.Header().Set("Expires", "Fri, 01 Jan 1980 00:00:00 GMT")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, must-revalidate")
	w.Header().Set("Content-Type", content_type)
	w.WriteHeader(http.StatusOK)
}

func CreateFirstPKTLine(service string) string {
	packet := fmt.Sprintf("# service=git-%s\n", service)

	prefix := strconv.FormatInt(int64(len(packet)+4), 16)
	if len(prefix)%4 != 0 {
		prefix = strings.Repeat("0", 4-len(prefix)%4) + prefix
	}
	magicMarker := "0000"
	return prefix + packet + magicMarker
}

func GetDefaultReposPath() string {
	rPath, _ := os.Getwd()
	return rPath
}

func IsExistingRepository(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func UserPath(userName string) string {
	return filepath.Join(config.ReposRootPath, strings.ToLower(userName))
}

func RepoPath(userName, repoName string) string {
	return filepath.Join(UserPath(userName), FormatRepoName(repoName))
}

func FormatRepoName(repoName string) string {
	var r string
	if strings.HasSuffix(repoName, ".git") {
		r = strings.ToLower(repoName)
	} else {
		r = strings.ToLower(repoName) + ".git"
	}
	return r
}

func GetDefaultGitPath() string {
	return "/usr/bin/git"
}
