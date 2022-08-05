package models

import (
	"bytes"
	"encoding/json"
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
