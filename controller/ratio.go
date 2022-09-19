package controller

import (
	"awesomeProject1/models"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type response struct {
	CompanyName       string      `json:"company_name"`
	Npwp              interface{} `json:"npwp"`
	BankDebtToEquity  *float64    `json:"bank_debt_to_equity"`
	Capitalisation    *float64    `json:"capitalisation"`
	GrossProfitMargin *float64    `json:"gross_profit_margin"`
	CurrentRatio      *float64    `json:"current_ratio"`
}

type PbkDumps struct {
	CompanyName string      `json:"company_name"`
	Npwp        interface{} `json:"npwp"`
	ApiKey      string      `json:"api_key"`
}

type Payload struct {
	CompanyName string      `json:"company_name"`
	Npwp        interface{} `json:"npwp"`
}

func (request PbkDumps) validatedev() url.Values {

	errs := url.Values{}

	if request.CompanyName == "" {
		errs.Add("company_name", "Key is required")
	}

	if request.Npwp == "" {
		errs.Add("npwp", "Key is required")
	}

	return errs
}

func (request PbkDumps) validateprod() url.Values {

	errs := url.Values{}

	if request.CompanyName == "" {
		errs.Add("company_name", "Key is required")
	}

	if request.Npwp == "" {
		errs.Add("npwp", "Key is required")
	}

	if request.ApiKey == "" {
		errs.Add("api_key", "Key is required")
	}

	if request.ApiKey != os.Getenv("API_KEY") {
		errs.Add("api_key", "api_key is wrong")
	}

	return errs
}

func PostCompany(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validateprod(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	Payload := Payload{
		CompanyName: PbkDumps.CompanyName,
		Npwp:        PbkDumps.Npwp,
	}

	data := models.GetCompany(Payload.CompanyName, Payload.Npwp)

	f := func(s float64) *float64 {
		return &s
	}

	if data.CompanyName == "No Data" {
		err := map[string]interface{}{"Error": "No company found"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	res := response{
		CompanyName:       data.CompanyName,
		Npwp:              PbkDumps.Npwp.(string),
		BankDebtToEquity:  data.BankDebtToEquity,
		Capitalisation:    data.Capitalisation,
		GrossProfitMargin: data.GrossProfitMargin,
		CurrentRatio:      data.CurrentRatio,
	}

	_ = json.NewEncoder(w).Encode(res)
}

func PostCompanyStaging(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validateprod(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	Payload := Payload{
		CompanyName: PbkDumps.CompanyName,
		Npwp:        PbkDumps.Npwp,
	}

	data, status := models.GetCompanyStaging(Payload.CompanyName, Payload.Npwp)

	f := func(s float64) *float64 {
		return &s
	}

	if data.CompanyName == "No Data" {
		err := map[string]interface{}{"Error": "No company found"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if status == "Reject" {
		err := map[string]interface{}{"Error": "No company found"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if data.Capitalisation == f(0) && data.GrossProfitMargin == f(0) && data.BankDebtToEquity == f(0) && data.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	res := response{
		CompanyName:       data.CompanyName,
		Npwp:              PbkDumps.Npwp.(string),
		BankDebtToEquity:  data.BankDebtToEquity,
		Capitalisation:    data.Capitalisation,
		GrossProfitMargin: data.GrossProfitMargin,
		CurrentRatio:      data.CurrentRatio,
	}

	_ = json.NewEncoder(w).Encode(res)
}

func ConnectionTest(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validateprod(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	Payload := Payload{
		CompanyName: PbkDumps.CompanyName,
		Npwp:        PbkDumps.Npwp,
	}

	status := models.TestConnection(Payload.Npwp)
	_ = json.NewEncoder(w).Encode(status)
}

func PostCompanyDev(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validatedev(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	finratres := models.FetchFinratio(PbkDumps.CompanyName, PbkDumps.Npwp)
	f := func(s float64) *float64 {
		return &s
	}

	if finratres.CompanyName == "No Data" {
		err := map[string]interface{}{"Error": "No company found"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if finratres.Capitalisation == f(0) && finratres.GrossProfitMargin == f(0) && finratres.BankDebtToEquity == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if finratres.Capitalisation == f(0) && finratres.BankDebtToEquity == f(0) && finratres.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if finratres.Capitalisation == f(0) && finratres.GrossProfitMargin == f(0) && finratres.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if finratres.GrossProfitMargin == f(0) && finratres.BankDebtToEquity == f(0) && finratres.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	} else if finratres.Capitalisation == f(0) && finratres.GrossProfitMargin == f(0) && finratres.BankDebtToEquity == f(0) && finratres.CurrentRatio == f(0) {
		err := map[string]interface{}{"Error": "Variable not sufficient"}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	res := response{
		CompanyName:       finratres.CompanyName,
		Npwp:              PbkDumps.Npwp,
		BankDebtToEquity:  finratres.BankDebtToEquity,
		Capitalisation:    finratres.Capitalisation,
		GrossProfitMargin: finratres.GrossProfitMargin,
		CurrentRatio:      finratres.CurrentRatio,
	}

	_ = json.NewEncoder(w).Encode(res)

}
