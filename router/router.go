package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there!")
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

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	// mux.HandleFunc("/expenses", expensesHandler)
	// mux.HandleFunc("/incomes", incomesHandler)
	// mux.HandleFunc("/resources/time-series/expenses", timeSeriesExpensesHandler)
	// mux.HandleFunc("/resources/time-series/incomes", timeSeriesIncomesHandler)
	// mux.HandleFunc("/resources/time-series/savings", timeSeriesSavingsHandler)
	// mux.HandleFunc("/resources/expenses/percentile/", resourcesExpensesPercentileHandler)
	return router
}
