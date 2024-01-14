package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samuelralmeida/investiment-calc/entity"
)

type InvestimentRepositoryInterface interface {
	ListInvestiment(ctx context.Context) (*[]entity.Investiment, error)
	SaveInvestiment(ctx context.Context, investiment *entity.Investiment) error
	SaveInvestimentCheckpoints(ctx context.Context, checkpoints *[]entity.Checkpoint) error
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
	investiment.ID = uuid.NewString()
	investiment.Date = time.Now()
	return s.Repository.SaveInvestiment(ctx, investiment)
}

func (s *service) CreateInvestmentCheckpoint(ctx context.Context, checkpoints *[]entity.Checkpoint) error {
	return s.Repository.SaveInvestimentCheckpoints(ctx, checkpoints)
}
