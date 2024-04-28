package sqlite

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

type jsonField []string

func (jf *jsonField) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	data, ok := value.(string)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	err := json.Unmarshal([]byte(data), &jf)
	if err != nil {
		return fmt.Errorf("scan json field: %w", err)
	}

	return nil
}

func (jf jsonField) Value() (driver.Value, error) {
	b := new(strings.Builder)
	err := json.NewEncoder(b).Encode(jf)
	return b.String(), err
}

func sliceStringToJsonField(input []string) (driver.Value, error) {
	if len(input) == 0 {
		return nil, nil
	}

	var resp []string
	for _, value := range input {
		str := strings.TrimSpace(value)
		if str == "" {
			continue
		}
		resp = append(resp, str)
	}

	b := new(strings.Builder)
	err := json.NewEncoder(b).Encode(resp)
	return b.String(), err
}
