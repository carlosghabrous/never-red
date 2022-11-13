package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (handler *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

type WorldHandler struct{}

func (handler *WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}
	http.Handle("/hello", &helloHandler)
	http.Handle("/world", &worldHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server %e", err)
	}
}
