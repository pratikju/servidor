package main

import (
	"flag"
	"path/filepath"
)

type Config struct {
	Port         string
	Hostname     string
	GitPath      string
	RepoRootPath string
}

var (
	port     = flag.String("p", "8000", "Port on which git server will listen")
	hostName = flag.String("b", "0.0.0.0", "Hostname to be used")
	repo     = flag.String("r", GetDefaultRepoPath(), "Set the path where repositories will be saved, Just mention the base directory(\"repos\" directory will be automatically created inside it).")
	git      = flag.String("g", "/usr/bin/git", "Mention the gitPath if its different in the system")
	config   Config
)

func main() {
	flag.Parse()

	config = Config{Port: *port, Hostname: *hostName, RepoRootPath: filepath.Join(*repo, "repos"), GitPath: *git}
	GitServer()
}
