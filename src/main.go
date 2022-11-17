package main

import (
	"log"
	"net/http"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()
	mux.HandleFunc("/import", dataImport)

	serverAddr := "127.0.0.1" + ":" + "8080"
	server := http.Server{Addr: serverAddr, Handler: mux}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server %e", err)
	}
}
