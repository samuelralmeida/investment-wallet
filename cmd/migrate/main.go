package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/samuelralmeida/investment-wallet/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db, err := database.NewSqliteConnection()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS funds (
			id         TEXT PRIMARY KEY,
			name       TEXT NOT NULL,
			cnpj       TEXT NOT NULL,
			box        TEXT NOT NULL,
			flavor     TEXT NOT NULL,
			bank       TEXT NOT NULL,
			min_value  REAL NOT NULL,
			deleted_at DATETIME NULL
		);
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS investments (
			id            TEXT NOT NULL,
			fund_id       TEXT NOT NULL,
			date          DATE NOT NULL,
			value         REAL NOT NULL,
			wallet        TEXT NOT NULL,
			deleted_at	  DATETIME NULL,
			FOREIGN KEY (fund_id) REFERENCES funds(id)
		);
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS checkpoints (
			id            TEXT NOT NULL,
			fund_id       TEXT NOT NULL,
			date          DATE NOT NULL,
			value         REAL NOT NULL,
			wallet        TEXT NOT NULL,
			deleted_at	  DATETIME NULL,
			FOREIGN KEY (fund_id) REFERENCES funds(id)
		);
	`)

	if err != nil {
		panic(err)
	}
}
