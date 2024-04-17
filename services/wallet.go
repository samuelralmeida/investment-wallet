package services

import (
	"context"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type IWalletRepository interface {
	SelectWallets(ctx context.Context) ([]entity.Wallet, error)
	InsertWallet(ctx context.Context, wallet *entity.Wallet) error
}

func (s *Services) ListWallets(ctx context.Context) ([]entity.Wallet, error) {
	return s.Repository.SelectWallets(ctx)
}

func (s *Services) SaveWallet(ctx context.Context, wallet *entity.Wallet) error {
	return s.Repository.InsertWallet(ctx, wallet)
}
