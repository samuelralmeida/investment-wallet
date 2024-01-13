package entity

import "time"

type Investiment struct {
	ID       string
	Date     time.Time
	Box      string
	Category string
	Name     string
	Cnpj     string
	Bank     string
	Amount   float64
	Pocket   string
	DeleteAt *time.Time
}

type InvestimentCheckpoint struct {
	ID            int
	InvestimentID string
	Amount        float64
	DeleteAt      *time.Time
}
