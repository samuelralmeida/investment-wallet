package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/samuelralmeida/investiment-calc/entity"
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
		return fmt.Errorf("inser investiment: %w", err)
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

func (r *repository) ListInvestiment(ctx context.Context) (*[]entity.Investiment, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		"SELECT i.id, i.date, i.box, i.category, i.name, i.cnpj, i.bank, i.amount, i.wallet, i.deleted_at FROM investments i",
	)
	if err != nil {
		return nil, fmt.Errorf("select investments: %w", err)
	}

	investiments := []entity.Investiment{}

	for rows.Next() {
		var investiment entity.Investiment
		err := rows.Scan(
			&investiment.ID, &investiment.Date, &investiment.Box, &investiment.Category,
			&investiment.Name, &investiment.Cnpj, &investiment.Bank, &investiment.Amount,
			&investiment.Wallet, &investiment.DeleteAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan investiment row: %w", err)
		}

		investiments = append(investiments, investiment)
	}

	return &investiments, nil
}

func (r *repository) SaveInvestimentCheckpoints(ctx context.Context, checkpoints *[]entity.Checkpoint) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO checkpoints (investiment_id, date, value) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}

	for _, checkpoint := range *checkpoints {
		_, err = stmt.Exec(checkpoint.InvestimentID, checkpoint.Date, checkpoint.Value)
		if err != nil {
			return fmt.Errorf("insert checkpoint: %w", err)
		}
	}

	return tx.Commit()
}

func (r *repository) ListInestimentsWithCheckpoint(ctx context.Context, wallet string) (*entity.Shares, error) {

	rows, err := r.DB.QueryContext(ctx, `
			select i.id, i.date, i.box, i.category, i.name, i.cnpj, i.bank, i.amount, i.wallet, lc.checkpoint_date, lc.value
			from investiments i
			left join (
				SELECT investiment_id, MAX(date) AS checkpoint_date, value
				FROM checkpoints
				WHERE deleted_at is null
				GROUP BY investiment_id
			) lc on i.id = lc.investiment_id
			where i.wallet = ?
		`, wallet)

	if err != nil {
		return nil, fmt.Errorf("select checkpoints: %w", err)
	}

	shares := entity.Shares{}

	for rows.Next() {
		var share entity.Share
		var checkpointDateTemp *string
		err := rows.Scan(
			&share.ID, &share.Date, &share.Box, &share.Category, &share.Name, &share.Cnpj, &share.Bank,
			&share.Amount, &share.Wallet, &checkpointDateTemp, &share.CheckpointValue,
		)

		if checkpointDateTemp != nil {
			temp, err := time.Parse("2006-01-02", *checkpointDateTemp)
			if err != nil {
				log.Println("parse checkout time:", err)
				continue
			}

			share.CheckpointDate = &temp
		}

		if err != nil {
			return nil, fmt.Errorf("scan investimetn with checkpoint row: %w", err)
		}

		shares = append(shares, share)
	}

	return &shares, nil
}
