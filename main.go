package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port                string
	Hostname            string
	GitPath             string
	ReposRootPath       string
	AuthEnabled         bool
	PasswdFilePath      string
	SSLEnabled          bool
	RestrictReceivePack bool
	RestrictUploadPack  bool
}

var (
	port         = flag.String("p", "8000", "Port on which servidor will listen")
	hostName     = flag.String("b", "0.0.0.0", "Hostname to be used")
	repo         = flag.String("r", GetDefaultReposPath(), "Set the path where repositories will be saved, Just mention the base path(\"repos\" directory will be automatically created inside it)")
	gitPath      = flag.String("g", GetDefaultGitPath(), "Mention the gitPath if its different on hosting machine")
	passwdFile   = flag.String("c", "", "Set the path from where the password file is to be read(to be set whenever -a flag is used)")
	auth         = flag.Bool("a", false, "Enable basic authentication for all http operations")
	ssl          = flag.Bool("s", false, "Enable tls connection")
	restrictPush = flag.Bool("R", false, "Set Whether ReceivePack(push operation) will be restricted")
	restrictPull = flag.Bool("U", false, "Set Whether UploadPack(clone, pull, fetch operations) will be restricted")
	config       Config
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
		GitPath: *gitPath, AuthEnabled: *auth, PasswdFilePath: *passwdFile, SSLEnabled: *ssl,
		RestrictReceivePack: *restrictPush, RestrictUploadPack: *restrictPull}
	GitServer()
}
