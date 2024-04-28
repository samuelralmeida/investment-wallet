package services

import (
	"context"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type IFundRepository interface {
	SelectFunds(ctx context.Context) ([]entity.Fund, error)
	InsertFund(ctx context.Context, fund *entity.Fund, subCategoryID int) error
}

func (s *Services) ListFunds(ctx context.Context) ([]entity.Fund, error) {
	return s.Repository.SelectFunds(ctx)
}

func (s *Services) SaveFund(ctx context.Context, fund *entity.Fund, subCategoryID int) error {
	return s.Repository.InsertFund(ctx, fund, subCategoryID)
}
