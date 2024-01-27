package entity

import "time"

type Checkpoint struct {
	ID       string
	FundID   string
	Date     time.Time
	Value    float64
	Wallet   string
	DeleteAt *time.Time
}
