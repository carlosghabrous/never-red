package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

func world(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World!")
}

// type Config map[string]string

// var neverRedConfig Config

// func init() {
// 	neverRedConfig, err := godotenv.Read()
// 	if err != nil {
// 		log.Fatal("Could not read .env file")
// 	}
// 	// fmt.Println(neverRedConfig["NEVER_RED_HOST"], neverRedConfig["NEVER_RED_PORT"])
// }

// TODO: load configuration
// TODO: have the command line option of starting a TLS server?

func main() {
	serverAddr := "127.0.0.1" + ":" + "8080"
	server := http.Server{Addr: serverAddr, Handler: nil}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server %e", err)
	}
}
