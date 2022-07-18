package server

import (
	"log"
	"net/http"

	"github.com/carlosghabrous/never-red/router"

	"github.com/joho/godotenv"
)

// Stores the environment variables stored in the .env file
var neverRedEnv map[string]string

func loadDotEnv() {
	var err error
	neverRedEnv, err = godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
}

func Start() error {
	loadDotEnv()
	neverRedPort, found := neverRedEnv["NEVER_RED_PORT"]

	if !found {
		log.Fatal("Variable 'NEVER_RED_PORT' not defined in .env file")
	}

	log.Printf("Initiating server on port %s", neverRedPort)
	return http.ListenAndServe(":"+neverRedPort, router.Router())
}
