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
	r.HandleFunc("/", rootHandler).Methods("GET")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Just testing the server\n"))
}
