package main

import (
	"flag"
)

type Config struct {
	Port     string
	Hostname string
}

var (
	port     = flag.String("p", "8000", "Port on which git server will listen")
	hostName = flag.String("b", "0.0.0.0", "Hostname to be used")
	config   Config
)

func main() {
	flag.Parse()
	config = Config{Port: *port, Hostname: *hostName}
	GitServer()
}
