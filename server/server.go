package server

import (
	"fmt"
	"log"
	"net/http"
)

var mux *http.ServeMux

// swagger:route GET /hello App says hello
//
// Makes the app say hello
//
// responses:
//	200: OK
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there!\n")
}

func csvImportHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "csv import!\n")
}

func expensesHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "expenses handler!\n")
}

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

func loadHandlers() {
	mux = http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/csv-import", csvImportHandler)
	mux.HandleFunc("/expenses", expensesHandler)
	mux.HandleFunc("/incomes", incomesHandler)
	mux.HandleFunc("/resources/time-series/expenses", timeSeriesExpensesHandler)
	mux.HandleFunc("/resources/time-series/incomes", timeSeriesIncomesHandler)
	mux.HandleFunc("/resources/time-series/savings", timeSeriesSavingsHandler)
	mux.HandleFunc("/resources/expenses/percentile/", resourcesExpensesPercentileHandler)
}

func Start() error {
	loadHandlers()

	log.Println("Listening...")
	return http.ListenAndServe(":9000", mux)
}
