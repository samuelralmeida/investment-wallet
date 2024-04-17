package sqlite

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/investment-wallet/entity"
)

func (r *Repository) SelectWallets(ctx context.Context) ([]entity.Wallet, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM wallets")
	if err != nil {
		return nil, fmt.Errorf("select wallets: %w", err)
	}

	wallets := []entity.Wallet{}
	for rows.Next() {
		var wallet entity.Wallet
		err := rows.Scan(&wallet.ID, &wallet.Name)
		if err != nil {
			return nil, fmt.Errorf("scan wallet row: %w", err)
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func (r *Repository) InsertWallet(ctx context.Context, wallet *entity.Wallet) error {
	result, err := r.db.ExecContext(ctx, "INSERT INTO wallets (name) VALUES (?)", wallet.Name)
	if err != nil {
		return fmt.Errorf("insert wallet: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("get last insert wallet id: %w", err)
	}

	wallet.ID = int(lastID)
	return nil
}
