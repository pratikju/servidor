package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GitServer() {
	host := fmt.Sprintf("%s:%s", config.Hostname, config.Port)
	log.Println("Starting git http server at", host)

	r := mux.NewRouter()
	attachHandler(r)
	if err := http.ListenAndServe(host, r); err != nil {
		log.Fatal(err)
	}
}

func attachHandler(r *mux.Router) {
	//git methods Handler
	r.HandleFunc(`/{user-name}/{repo-name:([a-zA-Z0-9\-\.\_]+)}/info/refs`, serviceHandler).Methods("GET")
	r.HandleFunc(`/{user-name}/{repo-name:([a-zA-Z0-9\-\.\_]+)}/git-upload-pack`, uploadPackHandler).Methods("POST")
	r.HandleFunc(`/{user-name}/{repo-name:([a-zA-Z0-9\-\.\_]+)}/git-receive-pack`, receivePackHandler).Methods("POST")

	//APIs handlers
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc(GetRepoCreateUrl(), repoCreateHandler).Methods("POST")
	r.HandleFunc(GetReposUrl(), repoIndexHandler).Methods("GET")
	r.HandleFunc(GetRepoUrl(), repoShowHandler).Methods("GET")
	r.HandleFunc(GetBranchesUrl(), branchIndexHandler).Methods("GET")
	r.HandleFunc(GetBranchUrl(), branchShowHandler).Methods("GET")
}
