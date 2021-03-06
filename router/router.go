package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

func Router() *mux.Router {
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
