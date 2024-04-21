package services

import (
	"context"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type ICategoryRepository interface {
	SelectCategories(ctx context.Context) ([]entity.Category, error)
}

func (s *Services) ListCategories(ctx context.Context) ([]entity.Category, error) {
	return s.Repository.SelectCategories(ctx)
}
