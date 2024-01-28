package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConnection() (*sql.DB, error) {
	env := strings.ToUpper(os.Getenv("ENV"))
	filename := os.Getenv(fmt.Sprintf("%s_SQLITE_FILENAME", env))

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
