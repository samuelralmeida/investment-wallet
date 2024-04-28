package sqlite

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/investment-wallet/entity"
)

func (r *Repository) SelectFunds(ctx context.Context) ([]entity.Fund, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM funds")
	if err != nil {
		return nil, fmt.Errorf("select funds: %w", err)
	}

	funds := []entity.Fund{}
	for rows.Next() {
		var fund entity.Fund
		err := rows.Scan(&fund.ID, &fund.Name)
		if err != nil {
			return nil, fmt.Errorf("scan fund row: %w", err)
		}
		funds = append(funds, fund)
	}

	return funds, nil
}

func (r *Repository) InsertFund(ctx context.Context, fund *entity.Fund, subCategoryID int) error {
	notes, err := sliceStringToJsonField(fund.Notes)
	if err != nil {
		return fmt.Errorf("convert notes to json: %w", err)
	}

	result, err := r.db.ExecContext(
		ctx,
		"INSERT INTO funds (name, bank, cnpj, min_value, notes, benchmark, subcategory_id ) VALUES (?, ?, ?, ?, ?, ?, ?)",
		fund.Name, fund.Bank, fund.Cnpj, fund.MinValue, notes, fund.Benchmark, subCategoryID,
	)
	if err != nil {
		return fmt.Errorf("insert fund: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("get last insert fund id: %w", err)
	}

	fund.ID = int(lastID)
	return nil
}
