package server

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there!\n")
}

// TODO: Could be done via map iteration for all handlers
func loadHandlers() {
	http.HandleFunc("/hello", helloHandler)
}

func Start() error {
	loadHandlers()

	return http.ListenAndServe(":9000", nil)
}
