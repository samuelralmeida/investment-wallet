package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const sqliteFileName = "investments.db"

func NewSqliteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteFileName)
	if err != nil {
		return nil, fmt.Errorf("open sqlite connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping sqlite connection: %w", err)
	}

	return db, nil
}
