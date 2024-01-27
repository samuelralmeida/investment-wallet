package entity

import (
	"time"
)

type Fund struct {
	ID        string
	Name      string
	Cnpj      string
	Box       string
	Flavor    string
	Bank      string
	MinValue  float64
	DeletedAt *time.Time
}

type Funds []Fund
