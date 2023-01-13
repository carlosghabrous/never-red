package server

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"
)

var testApp *App

func TestMain(m *testing.M) {
	var err error
	testApp, err = New()

	if err != nil {
		log.Fatal("Could not start App for test")
	}

	code := m.Run()

	//TODO: tear down
	os.Exit(code)
}

func TestDataImport(t *testing.T) {
	// Make request to upload CSV file
	// Query the DB and check that the record is there
	req := getMultipartRequest()
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetMovements(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/movements", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if respBody := response.Body.String(); respBody != "[]\n" {
		t.Errorf("Expected empty array; got %s\n", respBody)
	}
}

func TestGetMovement(t *testing.T) {
	// TODO: how to pass parameters to the query?
	// TODO: how to prepare the DB (aka, factories) for testing?
	req, _ := http.NewRequest(http.MethodGet, "/movement", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if respBody := response.Body.String(); respBody != "[]" {
		t.Errorf("Expected empty array; got %s\n", respBody)
	}
}

func TestGetNonExistingMovement(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/movement", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if respBody := response.Body.String(); respBody != "" {
		t.Errorf("Expected error message; got %s\n", respBody)
	}
}

// TODO: maybe in a testing utils file
func executeRequest(r *http.Request) (recorder *httptest.ResponseRecorder) {
	reqRecorder := httptest.NewRecorder()
	testApp.Router().ServeHTTP(reqRecorder, r)

	return reqRecorder
}

// TODO: maybe in a testing utils file
func checkResponseCode(t *testing.T, expected, got int) {
	if expected != got {
		t.Errorf("Expected response code %d; got %d\n", expected, got)
	}
}

func getMultipartRequest() *http.Request {
	cwd, _ := os.Getwd()
	dataFilePath := path.Join(cwd, "..", "..", "..", "data", "movements.csv")
	file, _ := os.Open(dataFilePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("csv", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", "/import", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}
