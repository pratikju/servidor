package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port           string
	Hostname       string
	GitPath        string
	ReposRootPath  string
	AuthEnabled    bool
	PasswdFilePath string
	SSLEnabled     bool
}

var (
	port       = flag.String("p", "8000", "Port on which git server will listen")
	hostName   = flag.String("b", "0.0.0.0", "Hostname to be used")
	repo       = flag.String("r", GetDefaultReposPath(), "Set the path where repositories will be saved, Just mention the base directory(\"repos\" directory will be automatically created inside it).")
	gitPath    = flag.String("g", GetDefaultGitPath(), "Mention the gitPath if its different in the system")
	passwdFile = flag.String("c", "", "Set the path from where the password file is to be read(to be set whenever -a flag is used)")
	auth       = flag.Bool("a", false, "Enable basic authentication for all http operations")
	ssl        = flag.Bool("s", false, "Allow ssl")
	config     Config
)

func main() {
	flag.Parse()
	if *auth {
		if *passwdFile == "" {
			log.Println("Improper usage")
			flag.Usage()
			return
		}
		if _, err := os.Open(*passwdFile); err != nil {
			log.Fatal(err)
		}
	}
	config = Config{Port: *port, Hostname: *hostName, ReposRootPath: filepath.Join(*repo, "repos"),
		GitPath: *gitPath, AuthEnabled: *auth, PasswdFilePath: *passwdFile, SSLEnabled: *ssl}
	GitServer()
}
