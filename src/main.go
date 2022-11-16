package main

import (
	"log"
	"net/http"
)

func dataImport(w http.ResponseWriter, r *http.Request) {
	if !validMethod("POST", r.Method, w) {
		return
	}
	//TODO: keep processing
}

func main() {
	serverAddr := "127.0.0.1" + ":" + "8080"
	server := http.Server{Addr: serverAddr, Handler: nil}

	http.HandleFunc("/import", dataImport)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server %e", err)
	}
}
