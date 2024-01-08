package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExplainAPIHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(explainAPI)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "AutoPrint API"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInvalidOSGetPrintersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/printers?os=invalidOS", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPrinters)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Invalid OS specified"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCorsHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(explainAPI)

	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("handler returned wrong CORS header: got %v want %v",
			rr.Header().Get("Access-Control-Allow-Origin"), "*")
	}
}

func TestPrintImageHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/print", strings.NewReader("image=test.jpg&copies=2&os=linux&printer=myPrinter"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(printImage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Printing image test.jpg 2 times on printer myPrinter"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
