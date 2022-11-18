package server

import "net/http"

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	theApi := &api{}

	mux := http.NewServeMux()
	http.HandleFunc("/import", dataImport)

	theApi.router = mux
	return theApi
}

func (a *api) Router() http.Handler {
	return a.router
}
