package main

import (
	"BookStore-API/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	var testApp app.App
	testApp.Initialize()

	recorder := httptest.NewRecorder()
	testApp.Server.ServeHTTP(recorder, req)

	return recorder
}

func Test_createBooks(t *testing.T) {

	requestBody := strings.NewReader(`{"title": "Buku Dongeng","author": "Guru TK"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/books", requestBody)
	request.Header.Add("Content-Type", "application/json")

	response := executeRequest(request)

	assert.Equal(t, 200, response.Code)

}

func Test_getBook(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/books/1", nil)
	request.Header.Add("Content-Type", "application/json")

	response := executeRequest(request)

	assert.Equal(t, 200, response.Code)
}

func Test_getBooks(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/books", nil)
	request.Header.Add("Content-Type", "application/json")

	response := executeRequest(request)

	assert.Equal(t, 200, response.Code)
}

func Test_updateBook(t *testing.T) {
	requestBody := strings.NewReader(`{"title": "Buku PKN","author": "Guru MI"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/books/3", requestBody)
	request.Header.Add("Content-Type", "application/json")

	response := executeRequest(request)

	assert.Equal(t, 200, response.Code)

}
func Test_deleteBook(t *testing.T) {

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/books/3", nil)
	request.Header.Add("Content-Type", "application/json")

	response := executeRequest(request)

	assert.Equal(t, 200, response.Code)
}
