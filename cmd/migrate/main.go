package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/samuelralmeida/investiment-calc/database"
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

	/*
	   _, err = db.Exec(`

	   	CREATE TABLE IF NOT EXISTS investiments (
	   		id         TEXT PRIMARY KEY,
	   		date       DATETIME NOT NULL,
	   		box        TEXT NOT NULL,
	   		category   TEXT NOT NULL,
	   		name       TEXT NOT NULL,
	   		cnpj       TEXT NOT NULL,
	   		bank       TEXT NOT NULL,
	   		amount     REAL NOT NULL,
	   		wallet     TEXT NOT NULL,
	   		deleted_at DATETIME NULL
	   	);

	   `)

	   	if err != nil {
	   		panic(err)
	   	}

	   _, err = db.Exec(`

	   	CREATE TABLE IF NOT EXISTS checkpoints (
	   		id             INTEGER PRIMARY KEY AUTOINCREMENT,
	   		investiment_id TEXT NOT NULL,
	   		date           DATETIME NOT NULL,
	   		value         REAL NOT NULL,
	   		deleted_at	   DATETIME NULL,
	   		FOREIGN KEY (investiment_id) REFERENCES investiments(id)
	   	);

	   `)

	   	if err != nil {
	   		panic(err)
	   	}

	   _, err = db.Exec(`

	   		INSERT INTO investiments (
	   			id, date, box,
	   			category, name, cnpj, bank, amount, wallet
	   		)
	   		VALUES(
	   			'1f15914a-a40e-4c44-a4d4-c1289b34bf5f', '2024-01-13 20:02:22.258427399 ', 'estabilidade',
	   			'baunilha', 'capitania premium 45', '20146294000171', 'btg', 30.5, 'principal'
	   		);
	   	`)

	   	if err != nil {
	   		panic(err)
	   	}

	   _, err = db.Exec(`

	   		INSERT INTO investiments (
	   			id, date, box,
	   			category, name, cnpj, bank, amount, wallet
	   		)
	   		VALUES(
	   			'd628eb7c-2107-40c7-8e04-aeb0b7605205', '2024-01-13 20:05:22.258427399 ', 'diversificação',
	   			'viés macro', 'kinea atlas', '29762315000158', 'btg', 31.00, 'principal'
	   		);
	   	`)

	   	if err != nil {
	   		panic(err)
	   	}
	*/
}
