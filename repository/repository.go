package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) SaveFund(ctx context.Context, fund *entity.Fund) error {
	_, err := r.DB.ExecContext(
		ctx,
		`
			INSERT INTO funds (id, name, cnpj, box, flavor, bank, min_value)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`,
		fund.ID, fund.Name, fund.Cnpj, fund.Box, fund.Flavor, fund.Bank, fund.MinValue,
	)

	if err != nil {
		return fmt.Errorf("insert fund: %w", err)
	}
	return nil
}

func (r *repository) SelectFunds(ctx context.Context) (*entity.Funds, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		"SELECT id, name, cnpj, box, flavor, bank, min_value FROM funds WHERE deleted_at IS NULL",
	)
	if err != nil {
		return nil, fmt.Errorf("select funds: %w", err)
	}

	funds := entity.Funds{}

	for rows.Next() {
		var fund entity.Fund
		err := rows.Scan(&fund.ID, &fund.Name, &fund.Cnpj, &fund.Box, &fund.Flavor, &fund.Bank, &fund.MinValue)

		if err != nil {
			return nil, fmt.Errorf("scan fund row: %w", err)
		}

		funds = append(funds, fund)
	}

	return &funds, nil
}

func (r *repository) SaveInvestment(ctx context.Context, investment *entity.Investment) error {
	_, err := r.DB.ExecContext(
		ctx,
		`
			INSERT INTO investments (id, fund_id, date, value, wallet)
			VALUES (?, ?, ?, ?, ?)
		`,
		investment.ID, investment.FundID, investment.Date.Format(time.DateOnly), investment.Value, investment.Wallet,
	)

	if err != nil {
		return fmt.Errorf("insert investment: %w", err)
	}
	return nil
}

func (r *repository) SaveCheckpoint(ctx context.Context, checkpoint *entity.Checkpoint) error {
	_, err := r.DB.ExecContext(
		ctx,
		`
			INSERT INTO checkpoints (id, fund_id, date, value, wallet)
			VALUES (?, ?, ?, ?, ?)
		`,
		checkpoint.ID, checkpoint.FundID, checkpoint.Date.Format(time.DateOnly), checkpoint.Value, checkpoint.Wallet,
	)
	if err != nil {
		return fmt.Errorf("insert checkpoint: %w", err)
	}

	return nil
}

func (r *repository) SelectInvestmentsByWallet(ctx context.Context, wallet string) (*entity.Investments, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		"SELECT id, fund_id, date, value, wallet FROM investments WHERE wallet = ? AND deleted_at IS NULL ORDER BY date ASC",
		wallet,
	)
	if err != nil {
		return nil, fmt.Errorf("select investments by wallet: %w", err)
	}

	investments := entity.Investments{}

	for rows.Next() {
		var investment entity.Investment
		err := rows.Scan(&investment.ID, &investment.FundID, &investment.Date, &investment.Value, &investment.Wallet)

		if err != nil {
			return nil, fmt.Errorf("scan investment row: %w", err)
		}

		investments = append(investments, investment)
	}

	return &investments, nil
}

func (r *repository) SelectFundsByIds(ctx context.Context, ids []string) (*entity.Funds, error) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(
		"SELECT id, name, cnpj, box, flavor, bank, min_value FROM funds WHERE deleted_at IS NULL AND id IN (%s)",
		strings.Join(placeholders, ", "),
	)

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select funds by id: %w", err)
	}

	funds := entity.Funds{}

	for rows.Next() {
		var fund entity.Fund
		err := rows.Scan(&fund.ID, &fund.Name, &fund.Cnpj, &fund.Box, &fund.Flavor, &fund.Bank, &fund.MinValue)

		if err != nil {
			return nil, fmt.Errorf("scan fund row: %w", err)
		}

		funds = append(funds, fund)
	}

	return &funds, nil
}

func (r *repository) SelectLastCheckpointByFundIDAndWallet(ctx context.Context, fundID string, wallet string) (*entity.Checkpoint, error) {
	var checkpoint entity.Checkpoint

	err := r.DB.QueryRowContext(
		ctx,
		"SELECT id, fund_id, date, value, wallet FROM checkpoints WHERE fund_id = ? AND wallet = ? AND deleted_at IS NULL ORDER BY date DESC LIMIT 1",
		fundID, wallet,
	).Scan(&checkpoint.ID, &checkpoint.FundID, &checkpoint.Date, &checkpoint.Value, &checkpoint.Wallet)

	if errors.Is(err, sql.ErrNoRows) {
		return &checkpoint, nil
	}

	if err != nil {
		return nil, fmt.Errorf("select checkpoint by fundID and wallet: %w", err)
	}

	return &checkpoint, nil
}
