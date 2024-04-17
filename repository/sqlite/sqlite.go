package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type Repository struct {
	db *sql.DB
}

func NewConnection(conf *entity.Conf) (*Repository, error) {
	db, err := sql.Open("sqlite3", conf.SQLITE.FILENAME)
	if err != nil {
		return nil, fmt.Errorf("open sqlite connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping sqlite connection: %w", err)
	}

	return &Repository{db: db}, nil
}
