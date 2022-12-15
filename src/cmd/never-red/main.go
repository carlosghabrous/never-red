package main

import (
	"log"
	"net/http"

	"github.com/carlosghabrous/never-red/src/pkg/server"
)

//TODO: get server data from environment file or similar
func main() {
	server, err := server.New()

	if err != nil {
		log.Fatal("Could not start app!")
	}

	log.Fatal(http.ListenAndServe(":"+server.Port(), server.Router()))
}
