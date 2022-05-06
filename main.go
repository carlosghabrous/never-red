package main

import (
	"log"

	"github.com/carlosghabrous/never-red/server"
)

func main() {

	if err := server.Start(); err != nil {
		log.Fatal("Could not start server!")
	}
}
