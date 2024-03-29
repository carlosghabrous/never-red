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

type neverRedEnv struct {
	dbHost  string
	dbUser  string
	dbPwd   string
	dbName  string
	dbPort  string
	appPort string
}

type App struct {
	router http.Handler // interface
	db     DbQuerier    // interface
	logger *log.Logger
	env    *neverRedEnv
}

// Variable representing the web app and its resources
var neverRedApp *App

// App constructor
func New() (*App, error) {
	neverRedApp = &App{}
	neverRedApp.initialize()

	return neverRedApp, nil
}

// initializes app
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

	app.db = dbDriver{driver: dbInstance}
	app.logger.Println("App correctly initialized!")
}

// initializes from the environment file
func initAppEnv() (*neverRedEnv, error) {
	cwd, _ := os.Getwd()
	log.Printf("current directory is %s\n", cwd)
	envMap, err := godotenv.Read("/app/.env")

	appEnv := &neverRedEnv{
		dbHost:  envMap["POSTGRES_HOST"],
		dbPort:  envMap["POSTGRES_PORT"],
		dbUser:  envMap["NEVER_RED_USER"],
		dbPwd:   envMap["NEVER_RED_PWD"],
		dbName:  envMap["NEVER_RED_DB_NAME"],
		appPort: envMap["NEVER_RED_PORT"]}
	return appEnv, err
}

func initAppRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/import", neverRedApp.dataImport)
	mux.HandleFunc("/movements", neverRedApp.getMovements)
	mux.HandleFunc("/movement", neverRedApp.getMovement)

	return mux
}

func initDb(env *neverRedEnv) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.dbHost, env.dbUser, env.dbPwd, env.dbName, env.dbPort, "disable")
	dbConn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't connect to DB %v", err)
	}

	return dbConn, nil
}

func (app *App) Router() http.Handler {
	return app.router
}

func (app *App) Port() string {
	app.logger.Printf("Listening on port %s\n", app.env.appPort)
	return app.env.appPort
}
