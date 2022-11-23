package main

import (
	"log"
	"net/http"

	"github.com/carlosghabrous/never-red/pkg/server"
)

//TODO: get server data from environment file or similar
func main() {
	server, err := server.New()
	if err != nil {
		log.Fatal("Could not start app!")
	}

	log.Fatal(http.ListenAndServe(":8080", server.Router()))
}