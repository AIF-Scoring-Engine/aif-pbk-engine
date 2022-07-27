package models

import (
	"awesomeProject1/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PbkDumps struct {
	CompanyName         string `json:"company_name"`
	Npwp                string `json:"npwp"`
	NGoodAfterBad       int64  `json:"n_good_after_bad"`
	OutstandingAmount   int64  `json:"outstanding_amount"`
	BankUmum            int64  `json:"bank_umum"`
	MobNewestFac        int64  `json:"mob_newest_fac"`
	Lknb                int64  `json:"lknb"`
	NActCntrIn1M        int64  `json:"n_act_cntr_in_1_m"`
	NoofClosedContracts int64  `json:"no_of_closed_contracts"`
}

type Finrats struct {
	Id                int64   `json:"id"`
	CompanyName       string  `json:"company_name"`
	Npwp              string  `json:"npwp"`
	BankDebtToEquity  float64 `json:"bank_debt_to_equity"`
	Capitalisation    float64 `json:"capitalisation"`
	GrossProfitMargin float64 `json:"gross_profit_margin"`
	CurrentRatio      float64 `json:"current_ratio"`
}

func DumpCompany(PBKDumps PbkDumps) *int64 {

	db := config.CreateConnection()

	sqlStatement := `INSERT INTO pbk_dumps (company_name, npwp, n_good_after_bad, outstanding_amount, bank_umum, 
                       mob_newest_fac, lknb, n_act_cntr_in_1_m, no_of_closed_contracts) 
                       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var id *int64

	err := db.QueryRow(sqlStatement, PBKDumps.CompanyName, PBKDumps.Npwp, PBKDumps.NGoodAfterBad,
		PBKDumps.OutstandingAmount, PBKDumps.BankUmum, PBKDumps.MobNewestFac, PBKDumps.Lknb,
		PBKDumps.NActCntrIn1M, PBKDumps.NoofClosedContracts).Scan(&id)

	if err != nil {
		log.Fatalf("Can't insert data. %v", err)
	}

	fmt.Printf("Inserted data with transaction_id: %v \n", id)

	_ = db.Close()

	return id
}

func FetchFinratio(CompanyNames string) Finrats {
	db := config.CreateConnection()

	var finrat Finrats

	sqlStatement := fmt.Sprintf(`SELECT * FROM financial_ratios WHERE company_name='%s'`, CompanyNames)

	err := db.QueryRow(sqlStatement).Scan(&finrat.Id, &finrat.CompanyName,
		&finrat.Npwp, &finrat.BankDebtToEquity, &finrat.Capitalisation, &finrat.GrossProfitMargin, &finrat.CurrentRatio)

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

func DuplicateValidation(CompanyNames string) *int64 {
	db := config.CreateConnection()

	var valid *int64

	sqlStatement := fmt.Sprintf(`select id FROM pbk_dumps WHERE company_name='%s'`, CompanyNames)

	err := db.QueryRow(sqlStatement).Scan(&valid)
	_ = db.Close()

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data Already Exists")
		return valid
	case nil:
		return valid
	default:
		log.Fatalf("Can't run query. %v", err)
	}

	return valid
}
