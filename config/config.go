package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection() *sql.DB {
	//err := godotenv.Load(".env")
	//
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	DbUser := os.Getenv("USER_LOCAL")
	DbPassword := os.Getenv("PASSWORD_LOCAL")
	DbName := os.Getenv("DATABASE_LOCAL")
	DbHost := os.Getenv("HOST_LOCAL")
	DbPort := os.Getenv("PORT_LOCAL")
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
