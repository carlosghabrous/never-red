package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// TODO: interface for db connection too?
// TODO : add a "global" logger to the App struct?

type neverRedEnv map[string]string
type App struct {
	router http.Handler // interface
	db     *sql.DB
	logger *log.Logger
	env    neverRedEnv
}

// Variable representing the web app and its resources
var app *App

func New() (*App, error) {
	app := &App{}
	app.initialize()

	return app, nil
}

func (app *App) initialize() {
	// logger
	app.logger = log.Default()

	// environment
	env, err := initAppEnv()
	if err != nil {
		app.logger.Fatal(err)
	}
	app.env = env

	// multiplexer
	app.router = initAppRouter()

	// database
	dbInstance, err := initDb(env)

	if err != nil {
		app.logger.Fatal(err)
	}

	app.db = dbInstance
	app.logger.Println("App correctly initialized!")
}

func initAppEnv() (neverRedEnv, error) {
	cwd, _ := os.Getwd()
	log.Printf("current directory is %s\n", cwd)
	appEnv, err := godotenv.Read(".env")
	return appEnv, err
}

func initAppRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/import", app.dataImport)
	mux.HandleFunc("/movements", app.getMovements)

	return mux
}

func initDb(env neverRedEnv) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env["POSTGRES_HOST"], env["NEVER_RED_USER"], env["NEVER_RED_PWD"], env["NEVER_RED_DB_NAME"], env["POSTGRES_PORT"], "disable")
	dbConn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't connect to DB %v\n", err)
	}

	return dbConn, nil
}

func (app *App) Router() http.Handler {
	return app.router
}

func (app *App) DB() *sql.DB {
	return app.db
}

func (app *App) Port() string {
	return app.env["NEVER_RED_PORT"]
}
