package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samuelralmeida/investiment-calc/entity"
)

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) ListInvestiment(ctx context.Context) (*[]entity.Investiment, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM investiments")
	if err != nil {
		return nil, fmt.Errorf("select investiments: %w", err)
	}

	investiments := []entity.Investiment{}

	for rows.Next() {
		var investiment entity.Investiment
		err := rows.Scan(
			&investiment.ID, &investiment.Date, &investiment.Box, &investiment.Category,
			&investiment.Name, &investiment.Cnpj, &investiment.Bank, &investiment.Amount,
			&investiment.Pocket, &investiment.DeleteAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan investiment row: %w", err)
		}

		investiments = append(investiments, investiment)
	}

	return &investiments, nil
}

func (r *repository) SaveInvestiment(ctx context.Context, investiment *entity.Investiment) error {
	panic("not implemented")
}

func (r *repository) SaveInvestimentCheckpoints(ctx context.Context, checkpoints *[]entity.InvestimentCheckpoint) error {
	panic("not implemented")
}
