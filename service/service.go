package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/samuelralmeida/investment-wallet/entity"
)

type InvestimentRepositoryInterface interface {
	SaveFund(ctx context.Context, fund *entity.Fund) error
	SelectFunds(ctx context.Context) (*entity.Funds, error)
	SaveInvestment(ctx context.Context, investment *entity.Investment) error
	SaveCheckpoint(ctx context.Context, checkpoint *entity.Checkpoint) error
	SelectFundsByIds(ctx context.Context, ids []string) (*entity.Funds, error)
	SelectInvestmentsByWallet(ctx context.Context, wallet string) (*entity.Investments, error)
	SelectLastCheckpointByFundIDAndWallet(ctx context.Context, fundID string, wallet string) (*entity.Checkpoint, error)
}

type service struct {
	Repository InvestimentRepositoryInterface
}

func New(repository InvestimentRepositoryInterface) *service {
	return &service{Repository: repository}
}

func (s *service) CreateInvestiment(ctx context.Context, investiment *entity.Investment) error {
	investiment.ID = uuid.NewString()
	return s.Repository.SaveInvestment(ctx, investiment)
}

func (s *service) CreateFund(ctx context.Context, fund *entity.Fund) error {
	fund.ID = uuid.NewString()
	return s.Repository.SaveFund(ctx, fund)
}

func (s *service) ListFunds(ctx context.Context) (*entity.Funds, error) {
	return s.Repository.SelectFunds(ctx)
}

func (s *service) CreateCheckpoint(ctx context.Context, checkpoint *entity.Checkpoint) error {
	checkpoint.ID = uuid.NewString()
	return s.Repository.SaveCheckpoint(ctx, checkpoint)
}

func (s *service) Wallet(ctx context.Context, wallet string) (*entity.Wallet, error) {
	investments, err := s.Repository.SelectInvestmentsByWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}

	fundsMap := make(map[string]entity.Investments)
	fundIds := []string{}
	for _, investment := range *investments {
		fundID := investment.FundID
		_, ok := fundsMap[fundID]
		if !ok {
			fundIds = append(fundIds, fundID)
		}
		fundsMap[fundID] = append(fundsMap[fundID], investment)
	}

	funds, err := s.Repository.SelectFundsByIds(ctx, fundIds)
	if err != nil {
		return nil, err
	}

	fundsDetail := make([]entity.FundDetail, len(*funds))

	for i, fund := range *funds {
		checkpoint, err := s.Repository.SelectLastCheckpointByFundIDAndWallet(ctx, fund.ID, wallet)
		if err != nil {
			return nil, err
		}

		fundsDetail[i] = entity.FundDetail{
			Fund:        fund,
			Investments: fundsMap[fund.ID],
			Checkpoint:  *checkpoint,
		}
	}

	return &entity.Wallet{
		Name:        wallet,
		FundsDetail: fundsDetail,
	}, nil

}

func (s *service) Calculate(ctx context.Context, wallet string) (*entity.Wallet, error) {
	return s.Wallet(ctx, wallet)
}
