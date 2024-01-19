package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samuelralmeida/investiment-calc/entity"
)

type InvestimentRepositoryInterface interface {
	ListInvestiment(ctx context.Context) (*[]entity.Investiment, error)
	SaveInvestiment(ctx context.Context, investiment *entity.Investiment) error
	SaveInvestimentCheckpoints(ctx context.Context, checkpoints *[]entity.Checkpoint) error
	ListInestimentsWithCheckpoint(ctx context.Context, wallet string) (*entity.Shares, error)
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

func (s *service) Calculate(ctx context.Context) (*entity.Wallet, error) {
	investimentsWithCheckpoint, err := s.Repository.ListInestimentsWithCheckpoint(ctx, "principal")
	if err != nil {
		return nil, fmt.Errorf("fetch investiments: %w", err)
	}

	funds := investimentsWithCheckpoint.Funds()

	return &entity.Wallet{
		Name:  "principal",
		Funds: *funds,
		Boxes: *funds.Boxes(),
	}, nil

	/*
			fmt.Println("TOTAL")
			fmt.Println(result(funds.TotalAmount(), funds.TotalValue()))
			fmt.Println("--//--")

			fmt.Println("FUNDOS")
			for _, fund := range *funds {
				fmt.Println(fund.Name)
				fmt.Println(result(fund.Amount, fund.Value))
				fmt.Println("")
			}
			fmt.Println("--//--")

			fmt.Println("BOX")
			boxes := funds.Boxes()
			for _, box := range *boxes {
				fmt.Println(box.Name)
				fmt.Println(result(box.Amount, box.Value))
				fmt.Printf("%% total amount: %.2f\n", (box.Amount/funds.TotalAmount())*100)
				fmt.Printf("%% total value: %.2f\n", (box.Value/funds.TotalValue())*100)
				fmt.Println("")
			}
			fmt.Println("")

			fmt.Println("CATEGORIA")
			categories := funds.Categories()
			for _, category := range *categories {
				fmt.Println(category.Box, "-", category.Name)
				fmt.Println(result(category.Amount, category.Value))
				fmt.Printf("%% total amount: %.2f\n", (category.Amount/funds.TotalAmount())*100)
				fmt.Printf("%% total value: %.2f\n", (category.Value/funds.TotalValue())*100)
				fmt.Printf("%% box amount: %.2f\n", (category.Amount/boxes.TotalAmount(category.Box))*100)
				fmt.Printf("%% box value: %.2f\n", (category.Value/boxes.TotalValue(category.Box))*100)
				fmt.Println("")
			}
			fmt.Println("")



				totalRegister := investimentsWithCheckpoint.ToatalRegister()
				boxRegister := investimentsWithCheckpoint.BoxRegister()
				categoryRegister := investimentsWithCheckpoint.CategoryRegister()

				fmt.Println("TOTAL")
				fmt.Println(totalRegister.Result())
				fmt.Println("")

				fmt.Println("BOX")
				for box, register := range boxRegister {
					fmt.Println(box)
					fmt.Println(register.Result())
				}
				fmt.Println("")

				fmt.Println("CATEGROY")
				for category, register := range categoryRegister {
					fmt.Println(category)
					fmt.Println(register.Result())
				}
				fmt.Println("")


		return nil
	*/
}

func result(amount, value float64) string {
	return fmt.Sprintf(
		"amount: %.2f | value: %.2f | diff: %.2f | ratio: %.2f",
		amount,
		value,
		value-amount,
		(value/amount-1)*100,
	)
}
