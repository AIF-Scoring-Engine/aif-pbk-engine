package models

import (
	"awesomeProject1/config"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Response struct {
	CompanyName       string  `json:"company_name"`
	Npwp              string  `json:"npwp"`
	BankDebtToEquity  float64 `json:"bank_debt_to_equity"`
	Capitalisation    float64 `json:"capitalisation"`
	GrossProfitMargin float64 `json:"gross_profit_margin"`
	CurrentRatio      float64 `json:"current_ratio"`
}

type Finrats struct {
	Id                string  `json:"id"`
	CompanyName       string  `json:"company_name"`
	Npwp              string  `json:"npwp"`
	BankDebtToEquity  float64 `json:"bank_debt_to_equity"`
	Capitalisation    float64 `json:"capitalisation"`
	GrossProfitMargin float64 `json:"gross_profit_margin"`
	CurrentRatio      float64 `json:"current_ratio"`
}

type Payload struct {
	CompanyName string `json:"company_name"`
	Npwp        string `json:"npwp"`
}

func GetCompany(Payload Payload) Response {

	var result Response

	data, err := json.Marshal(Payload)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(data)) //replace with investree endpoint

	var respo map[string]interface{}

	_ = json.NewDecoder(resp.Body).Decode(&respo)
	result.CompanyName = respo["url"].(string)
	result.Npwp = respo["url"].(string)
	//result.BankDebtToEquity = respo["url"].(float64)
	//result.GrossProfitMargin = respo["url"].(float64)
	//result.Capitalisation = respo["url"].(float64)
	//result.Capitalisation = respo["url"].(float64)
	//fmt.Println(result.CompanyName)

	return result
}

func FetchFinratio(CompanyNames string, Npwp string) Finrats {
	db := config.CreateConnection()

	var finrat Finrats

	sqlStatement := fmt.Sprintf(`SELECT bld_loan_number, bpd_company_name, npwp, 
coalesce(capitalisation,0) capitalisation,
coalesce(gross_profit_margins,0) gross_profit_margins,
coalesce(bank_debt_to_equity,0) bank_debt_to_equity,
coalesce(current_ratio,0) current_ratio
FROM data_pbk.fin_ratio_investree WHERE bpd_company_name='%s' and npwp='%s' limit 1`, CompanyNames, Npwp)

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
