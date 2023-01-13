package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const MaxFileBytesInMemory = 1024

func (app *App) dataImport(w http.ResponseWriter, r *http.Request) {
	if !validMethod(http.MethodPost, r.Method, w) {
		return
	}

	r.ParseMultipartForm(MaxFileBytesInMemory)
	fileHeader := r.MultipartForm.File["csv"][0]
	file, err := fileHeader.Open()

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	_, err = ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	//TODO: parse file and dump content to DB
}

func (app *App) getMovements(w http.ResponseWriter, r *http.Request) {
	if !validMethod(http.MethodGet, r.Method, w) {
		return
	}

	// TODO: extract filters from request

	movements, err := app.db.getMovements()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Something happened while fetching data")
		return
	}

	json.NewEncoder(w).Encode(movements)
}

func (app *App) getMovement(w http.ResponseWriter, r *http.Request) {
	if !validMethod(http.MethodGet, r.Method, w) {
		return
	}

	// TODO: extract int from request
	app.logger.Printf("query is %s", r.URL.Query())
	id := 1
	movement, err := app.db.getMovement(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Something happened while fetching data")
		return
	}

	json.NewEncoder(w).Encode(movement)
}
