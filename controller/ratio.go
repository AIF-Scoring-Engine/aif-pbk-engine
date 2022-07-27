package controller

import (
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type response2 struct {
	TransactionId       *int64  `json:"transaction_id"`
	CompanyName         string  `json:"company_name"`
	Npwp                string  `json:"npwp"`
	NGoodAfterBad       int64   `json:"n_good_after_bad"`
	OutstandingAmount   int64   `json:"outstanding_amount"`
	BankUmum            int64   `json:"bank_umum"`
	MobNewestFac        int64   `json:"mob_newest_fac"`
	Lknb                int64   `json:"lknb"`
	NActCntrIn1M        int64   `json:"n_act_cntr_in_1_m"`
	NoofClosedContracts int64   `json:"no_of_closed_contracts"`
	BankDebtToEquity    float64 `json:"bank_debt_to_equity"`
	Capitalisation      float64 `json:"capitalisation"`
	GrossProfitMargin   float64 `json:"gross_profit_margin"`
	CurrentRatio        float64 `json:"current_ratio"`
}

type response struct {
	CompanyName       string  `json:"company_name"`
	Npwp              string  `json:"npwp"`
	BankDebtToEquity  float64 `json:"bank_debt_to_equity"`
	Capitalisation    float64 `json:"capitalisation"`
	GrossProfitMargin float64 `json:"gross_profit_margin"`
	CurrentRatio      float64 `json:"current_ratio"`
}

type PbkDumps struct {
	CompanyName string `json:"company_name"`
	Npwp        string `json:"npwp"`
}

func (request *PbkDumps) validate() url.Values {
	errs := url.Values{}

	if request.CompanyName == "" {
		errs.Add("company_name", "Key is required")
	}

	if request.Npwp == "" {
		errs.Add("npwp", "Key is required")
	}

	return errs
}

// Phase1

func PostCompany(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	finratres := models.FetchFinratio(PbkDumps.CompanyName)

	if finratres.CompanyName == "No Data" {
		err := map[string]interface{}{"validationError": finratres.CompanyName}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	res := response{
		CompanyName:       finratres.CompanyName,
		Npwp:              finratres.Npwp,
		BankDebtToEquity:  finratres.BankDebtToEquity,
		Capitalisation:    finratres.Capitalisation,
		GrossProfitMargin: finratres.GrossProfitMargin,
		CurrentRatio:      finratres.CurrentRatio,
	}

	_ = json.NewEncoder(w).Encode(res)

}

//Phase2

func PostCompany2(w http.ResponseWriter, r *http.Request) {

	PbkDumps := &PbkDumps{}

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)
	if err != nil {
		log.Fatalf("Can't decode from request body.  %v", err)
	}

	if validErrs := PbkDumps.validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "appliciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	data := url.Values{
		"company_name": {PbkDumps.CompanyName},
		"npwp":         {PbkDumps.Npwp},
	}

	resp, err := http.PostForm("https://httpbin.org/post", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])

}

// Phase3

func PostCompany3(w http.ResponseWriter, r *http.Request) {

	var PbkDumps models.PbkDumps

	err := json.NewDecoder(r.Body).Decode(&PbkDumps)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	var validator *int64

	validator = models.DuplicateValidation(PbkDumps.CompanyName)

	var insertID *int64

	if validator == nil {
		insertID = models.DumpCompany(PbkDumps)
	} else {
		insertID = validator
	}

	finratres := models.FetchFinratio(PbkDumps.CompanyName)

	res := response2{
		TransactionId:       insertID,
		CompanyName:         finratres.CompanyName,
		Npwp:                finratres.Npwp,
		NGoodAfterBad:       PbkDumps.NGoodAfterBad,
		OutstandingAmount:   PbkDumps.OutstandingAmount,
		BankUmum:            PbkDumps.BankUmum,
		MobNewestFac:        PbkDumps.MobNewestFac,
		Lknb:                PbkDumps.Lknb,
		NActCntrIn1M:        PbkDumps.NActCntrIn1M,
		NoofClosedContracts: PbkDumps.NoofClosedContracts,
		BankDebtToEquity:    finratres.BankDebtToEquity,
		Capitalisation:      finratres.Capitalisation,
		GrossProfitMargin:   finratres.GrossProfitMargin,
		CurrentRatio:        finratres.CurrentRatio,
	}

	_ = json.NewEncoder(w).Encode(res)

}
