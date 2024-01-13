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
	Amount   int
	Pocket   string
}
