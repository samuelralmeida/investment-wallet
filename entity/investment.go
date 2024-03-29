package entity

import (
	"fmt"
	"strings"
	"time"
)

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
