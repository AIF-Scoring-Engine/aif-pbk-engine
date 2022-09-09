package models

import (
	"awesomeProject1/config"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	CompanyName       string      `json:"company_name"`
	Npwp              interface{} `json:"npwp"`
	BankDebtToEquity  float64     `json:"bank_debt_to_equity"`
	Capitalisation    float64     `json:"capitalisation"`
	GrossProfitMargin float64     `json:"gross_profit_margin"`
	CurrentRatio      float64     `json:"current_ratio"`
}

type Finrats struct {
	Id                string      `json:"id"`
	CompanyName       string      `json:"company_name"`
	Npwp              interface{} `json:"npwp"`
	BankDebtToEquity  float64     `json:"bank_debt_to_equity"`
	Capitalisation    float64     `json:"capitalisation"`
	GrossProfitMargin float64     `json:"gross_profit_margin"`
	CurrentRatio      float64     `json:"current_ratio"`
}

type Payload struct {
	CompanyName string      `json:"company_name"`
	Npwp        interface{} `json:"npwp"`
}

type Match struct {
	Status   string
	Message  string
	DataDate string
	BpdId    string
	Data     struct {
		Capitalisation    float64 `json:"capitalisation"`
		GrossProfitMargin float64 `json:"gross_profit_margins"`
		BankDebtToEquity  float64 `json:"bank_debt_to_equity"`
		CurrentRatio      float64 `json:"current_ratio"`
	}
}

func GetCompany(CompanyNames string, Npwp interface{}) Response {

	var result Response

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	postData := bytes.NewBuffer([]byte(fmt.Sprintf(`{"filters":{"no_npwp":"%s"},"measure_names":["capitalisation","gross_profit_margins","bank_debt_to_equity","current_ratio"]}`, Npwp)))
	req, err := http.NewRequest("POST", "https://dw.investree.tech/v1/data-extraction/borrower-info", postData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("%s", os.Getenv("TOKEN_PRODUCTION")))

	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	respo := &Match{}

	err = json.Unmarshal([]byte(body), respo)
	if err != nil {
		log.Fatal(err)
	}

	if respo.Message != "Data Found" {
		result.BankDebtToEquity = 0
		result.GrossProfitMargin = 0
		result.Capitalisation = 0
		result.CurrentRatio = 0
	} else if respo.Message == "Data Found" {
		result.BankDebtToEquity = respo.Data.BankDebtToEquity
		result.GrossProfitMargin = respo.Data.GrossProfitMargin
		result.Capitalisation = respo.Data.Capitalisation
		result.CurrentRatio = respo.Data.CurrentRatio
	}

	result.CompanyName = CompanyNames
	result.Npwp = Npwp
	return result
}

func GetCompanyStaging(CompanyNames string, Npwp interface{}) Response {

	var result Response

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	postData := bytes.NewBuffer([]byte(fmt.Sprintf(`{"filters":{"no_npwp":"%s"},"measure_names":["capitalisation","gross_profit_margins","bank_debt_to_equity","current_ratio"]}`, Npwp)))
	req, err := http.NewRequest("POST", "https://dw.investree.tech/v1/data-extraction/borrower-info", postData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("%s", os.Getenv("TOKEN_STAGING")))

	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return result
	}

	respo := &Match{}

	err = json.Unmarshal([]byte(body), respo)
	if err != nil {
		log.Fatal(err)
	}

	if respo.Message != "Data Found" {
		result.BankDebtToEquity = 0
		result.GrossProfitMargin = 0
		result.Capitalisation = 0
		result.CurrentRatio = 0
	} else if respo.Message == "Data Found" {
		result.BankDebtToEquity = respo.Data.BankDebtToEquity
		result.GrossProfitMargin = respo.Data.GrossProfitMargin
		result.Capitalisation = respo.Data.Capitalisation
		result.CurrentRatio = respo.Data.CurrentRatio
	}

	result.CompanyName = CompanyNames
	result.Npwp = Npwp
	return result
}

func FetchFinratio(CompanyNames string, Npwp interface{}) Finrats {
	db := config.CreateConnection()

	var finrat Finrats

	sqlStatement := fmt.Sprintf(`SELECT bld_loan_number, bpd_company_name, npwp, 
coalesce(capitalisation,0) capitalisation,
coalesce(gross_profit_margins,0) gross_profit_margins,
coalesce(bank_debt_to_equity,0) bank_debt_to_equity,
coalesce(current_ratio,0) current_ratio
FROM data_pbk.fin_ratio_investree WHERE bpd_company_name=lower('%s') and npwp='%v' limit 1`, CompanyNames, Npwp)

	err := db.QueryRow(sqlStatement).Scan(&finrat.Id, &finrat.CompanyName,
		&finrat.Npwp, &finrat.Capitalisation, &finrat.GrossProfitMargin, &finrat.BankDebtToEquity, &finrat.CurrentRatio)

	_ = db.Close()

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No Data..")
		finrat.CompanyName = "No Data"
		return finrat
	case nil:
		return finrat
	default:
		log.Fatalf("Can't run query. %v", err)
	}

	return finrat
}
