package service

import (
	"context"
	"fmt"

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

func (s *service) Recommendation(ctx context.Context, walletName string) (*entity.Wallet, error) {
	wallet, err := s.Wallet(ctx, walletName)
	if err != nil {
		return nil, err
	}

	recommendation := struct {
		Current float64
		Invest  float64
		Total   float64
		Resp    map[string]float64
	}{}

	recommendation.Current = wallet.TotalCurrentValue()
	recommendation.Invest = 3000
	recommendation.Total = recommendation.Current + recommendation.Invest
	recommendation.Resp = make(map[string]float64)

	rules := map[string]float64{
		"ESTABILIDADE":    0.225,
		"DIVERSIFICAÇÃO":  0.4,
		"VALORIZAÇÃO":     0.3,
		"ANTIFRAGILIDADE": 0.075,
	}

	boxes := wallet.Boxes()

	for _, box := range boxes {
		goalRule := rules[box.Name]
		goalInvest := recommendation.Total * goalRule
		recommendation.Resp[box.Name] = goalInvest - box.CurrentValue
		fmt.Println(box.Name, recommendation.Current, recommendation.Invest, recommendation.Current+recommendation.Invest, box.CurrentValue, goalRule, goalInvest, goalInvest-box.CurrentValue)
	}

	for _, box := range boxes {
		goalRule := rules[box.Name]
		goalInvest := recommendation.Total * goalRule
		// recommendation

		recommendation.Resp[box.Name] = goalInvest - box.CurrentValue
		fmt.Println(box.Name, recommendation.Current, recommendation.Invest, recommendation.Current+recommendation.Invest, box.CurrentValue, goalRule, goalInvest, goalInvest-box.CurrentValue)
	}

	fmt.Printf("%+v\n", recommendation)
	return nil, nil
}
