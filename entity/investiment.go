package entity

import (
	"fmt"
	"strings"
	"time"
)

type Investiment struct {
	ID       string
	Date     time.Time
	Box      string
	Category string
	Name     string
	Cnpj     string
	Bank     string
	Amount   float64
	Wallet   string
	DeleteAt *time.Time
}

type Checkpoint struct {
	ID            int
	InvestimentID string
	Date          time.Time
	Value         float64
	DeleteAt      *time.Time
}

type Investment struct {
	ID       string
	FundID   string
	Date     time.Time
	Value    float64
	Wallet   string
	DeleteAt *time.Time
}

func (i *Investment) FormatDate() string {
	return i.Date.Format("02/01/2006")
}

func (i *Investment) FormatValue() string {
	return strings.Replace(fmt.Sprintf("$%.2f", i.Value), ".", ",", 1)
}

type Investments []Investment

type FundDetail struct {
	Fund        Fund
	Investments Investments
}

type Wallet2 struct {
	Name        string
	FundsDetail []FundDetail
}
