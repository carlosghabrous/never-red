package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const MaxFileBytesInMemory = 1024

func dataImport(w http.ResponseWriter, r *http.Request) {
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
