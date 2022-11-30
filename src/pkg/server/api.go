package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
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
	app.initialize()

	return app, nil
}

func (app *App) initialize() {
	app.router = initAppRouter()
	dbInstance, err := initDb()

	if err != nil {
		log.Fatal(err)
	}

	app.db = dbInstance
}

func initAppRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/import", app.dataImport)
	mux.HandleFunc("/movements", app.getMovements)

	return mux
}

func initDb() (*sql.DB, error) {
	//TODO: get DB credentials from .env file
	const (
		user     = "never_red_user"
		password = "super_secret_never_red_pwd"
		dbname   = "never_red"
		host     = "localhost"
		port     = 5434
		sslmode  = "disable"
	)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", user, password, dbname, host, port, sslmode)
	dbConn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func (app *App) Router() http.Handler {
	return app.router
}

func (app *App) DB() *sql.DB {
	return app.db
}
