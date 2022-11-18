package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const MaxFileBytesInMemory = 1024

func (a *api) dataImport(w http.ResponseWriter, r *http.Request) {
	if !validMethod("POST", r.Method, w) {
		return
	}

	r.ParseMultipartForm(MaxFileBytesInMemory)
	fileHeader := r.MultipartForm.File["csv"][0]
	file, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		// return?
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		// return?
	}
	fmt.Fprintln(w, data)

	//TODO: parse file and dump content to DB
}

func getMovements(w http.ResponseWriter, r *http.Request) {
	if !validMethod("GET", r.Method, w) {
		return
	}

	movements := json.Number(1) // this should come from the DB...
	var err error = nil         // ...and the error as well

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Something happened while fetching data")
		return
	}

	json.NewEncoder(w).Encode(movements)
}
