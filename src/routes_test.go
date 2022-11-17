package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataImport(t *testing.T) {
	// Make request to upload CSV file
	// Query the DB and check that the record is there
	req, _ := http.NewRequest("POST", "/import", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected and empty array. Got %s\n", body)
	}
}

func executeRequest(r *http.Request) (recorder *httptest.ResponseRecorder) {
	reqRecorder := httptest.NewRecorder()
	mux.ServeHTTP(reqRecorder, r)

	return reqRecorder
}

func checkResponseCode(t *testing.T, expected, got int) {
	if expected != got {
		t.Errorf("Expected response code %d; got %d\n", expected, got)
	}
}
