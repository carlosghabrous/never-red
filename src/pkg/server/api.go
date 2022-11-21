package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

// TODO: interface for db connection too?
// TODO : add a "global" logger to the App struct?
type App struct {
	router http.Handler // interface
	db     *sql.DB
}

// Variable representing the web app and its resources
var app *App

func New() (*App, error) {
	app := &App{}
	app.router = initAppRouter()

	var err error
	app.db, err = initDb()

	if err != nil {
		return nil, fmt.Errorf("Could not connect to the DB: %w", err)
	}
	return app, nil
}

func initAppRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/import", app.dataImport)
	mux.HandleFunc("/movements", app.getMovements)

	return mux
}

func initDb() (*sql.DB, error) {
	//TODO: get DB credentials from somewhere
	return nil, nil
}

func (app *App) Router() http.Handler {
	return app.router
}
