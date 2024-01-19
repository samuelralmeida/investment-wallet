package entity

import (
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
