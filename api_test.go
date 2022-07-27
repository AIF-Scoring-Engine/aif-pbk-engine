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
    "company_name": "ditskih.co",
    "npwp":"1111",
    "n_good_after_bad":72,
    "outstanding_amount":1000000000,
    "bank_umum":12,
    "mob_newest_fac":123,
    "lknb":12,
    "n_act_cntr_in_1_m":123,
    "no_of_closed_contracts":123}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompany)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"company_name":"ditskih.co",` +
		`"npwp":"1111",` +
		`"bank_debt_to_equity":7.2,` +
		`"capitalisation":3.32,` +
		`"gross_profit_margin":1.23,` +
		`"current_ratio":0.65}`

	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostNoData(t *testing.T) {

	var jsonTest = []byte(`{
    "company_name": "ditskiih.co",
    "npwp":"1111",
    "n_good_after_bad":72,
    "outstanding_amount":1000000000,
    "bank_umum":12,
    "mob_newest_fac":123,
    "lknb":12,
    "n_act_cntr_in_1_m":123,
    "no_of_closed_contracts":123}`)

	req, err := http.NewRequest("POST", "localhost:8080/api/post/company", bytes.NewBuffer(jsonTest))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostCompany)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"validationError":"No Data"}`

	if strings.TrimRight(rr.Body.String(), "\n") != strings.TrimRight(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
