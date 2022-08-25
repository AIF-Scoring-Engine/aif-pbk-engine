package main

import (
	"awesomeProject1/controller"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostOK(t *testing.T) {

	var jsonTest = []byte(`{
    "company_name": "merpati wahana raya",
    "npwp":"21831094022000",
    "api_key":"tj3hnFdcc31U"
}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/dev/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompanyDev)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"company_name":"merpati wahana raya","npwp":"21831094022000","bank_debt_to_equity":1.24137389744408,"capitalisation":3.18445164329251,"gross_profit_margin":0.445269105232753,"current_ratio":1.61141012964412}`
	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostWrongApiKey(t *testing.T) {

	var jsonTest = []byte(`{
    "company_name": "merpati wahana raya",
    "npwp":"21831094022000",
    "api_key":"tj3hnFcc31U"
}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/dev/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompanyDev)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"validationError":{"api_key":["api_key is wrong"]}}`

	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostNoNpwp(t *testing.T) {

	var jsonTest = []byte(`{
    "company_name": "merpati wahana raya",
    "npwp":"",
    "api_key":"tj3hnFdcc31U"
}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/dev/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompanyDev)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"validationError":{"npwp":["Key is required"]}}`

	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostNoCompanyName(t *testing.T) {

	var jsonTest = []byte(`{
    "company_name": "",
    "npwp":"21831094022000",
    "api_key":"tj3hnFdcc31U"
}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/dev/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompanyDev)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"validationError":{"company_name":["Key is required"]}}`

	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
