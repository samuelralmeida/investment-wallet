package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const sqliteFileName = "investments%s.db"

func NewSqliteConnection() (*sql.DB, error) {
	env := os.Getenv("ENV")
	filename := fmt.Sprintf(sqliteFileName, env)

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, fmt.Errorf("open sqlite connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping sqlite connection: %w", err)
	}

	return db, nil
}
