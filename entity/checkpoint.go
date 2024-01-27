package entity

import (
	"fmt"
	"strings"
	"time"
)

type Checkpoint struct {
	ID       string
	FundID   string
	Date     time.Time
	Value    float64
	Wallet   string
	DeleteAt *time.Time
}

func (c *Checkpoint) FormatDate() string {
	return c.Date.Format("02/01/2006")
}

func (c *Checkpoint) FormatValue() string {
	return strings.Replace(fmt.Sprintf("$%.2f", c.Value), ".", ",", 1)
}
