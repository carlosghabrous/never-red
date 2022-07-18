package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Stores the environment variables stored in the .env file
var neverRedEnv map[string]string

// swagger:route GET /hello Hello
//
// Makes the app say hello
//
// responses:
// 		200
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there!\n")
}

// swagger:route POST /csv-import Import CSV file
//
// Imports data from a CSV file in a predefined format
//
// responses:
// 		201: OK
func csvImportHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "csv import!\n")
}

// swagger:route GET expenses Expenses
//
// Queries the DB to get expenses' data
//
// responses:
// 		200: OK
func expensesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "expenses handler!\n")
}

// swagger:route GET incomes Incomes
//
// Queries the DB to get incomes' data
//
// responses:
//		200: OK
func incomesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "incomes handler!\n")
}

func timeSeriesExpensesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "timeSeriesExpenses handler!\n")
}

func timeSeriesIncomesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "timeSeriesIncomes handler!\n")
}

func timeSeriesSavingsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "timeSeriesSavings handler!\n")
}

func resourcesExpensesPercentileHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "resourcesExpensesPercentile handler!\n")
}

func loadDotEnv() {
	var err error
	neverRedEnv, err = godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
}

func loadHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	// mux.HandleFunc("/expenses", expensesHandler)
	// mux.HandleFunc("/incomes", incomesHandler)
	// mux.HandleFunc("/resources/time-series/expenses", timeSeriesExpensesHandler)
	// mux.HandleFunc("/resources/time-series/incomes", timeSeriesIncomesHandler)
	// mux.HandleFunc("/resources/time-series/savings", timeSeriesSavingsHandler)
	// mux.HandleFunc("/resources/expenses/percentile/", resourcesExpensesPercentileHandler)
	return router
}

func Start() error {
	loadDotEnv()
	neverRedPort, found := neverRedEnv["NEVER_RED_PORT"]

	if !found {
		log.Fatal("Variable 'NEVER_RED_PORT' not defined in .env file")
	}

	log.Printf("Initiating server on port %s", neverRedPort)
	return http.ListenAndServe(":"+neverRedPort, loadHandlers())
}
