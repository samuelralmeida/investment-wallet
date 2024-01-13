package service

import (
	"context"

	"github.com/samuelralmeida/investiment-calc/entity"
)

type InvestimentRepositoryInterface interface {
	ListInvestiment(ctx context.Context) (*[]entity.Investiment, error)
	SaveInvestiment(ctx context.Context, investiment *entity.Investiment) error
	SaveInvestimentCheckpoints(ctx context.Context, checkpoints *[]entity.InvestimentCheckpoint) error
}

type service struct {
	Repository InvestimentRepositoryInterface
}

func New(repository InvestimentRepositoryInterface) *service {
	return &service{Repository: repository}
}

func (s *service) ListInvestiments(ctx context.Context) (*[]entity.Investiment, error) {
	return s.Repository.ListInvestiment(ctx)
}

func (s *service) CreateInvestiment(ctx context.Context, investiment *entity.Investiment) error {
	panic("not implemented")
}

func (s *service) CreateInvestmentCheckpoint(ctx context.Context, investimentCheckpoints *[]entity.InvestimentCheckpoint) error {
	panic("not implemented")
}
