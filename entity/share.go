package entity

import "time"

// Share is each unit of a investiment fund
type Share struct {
	Investiment
	CheckpointDate  *time.Time
	CheckpointValue *float64
}

type Shares []Share

// Funds groups units of the same investiment fund
// When the share does not have a checkpoint, the checkpoint value is equal to the initial found value
func (ii *Shares) Funds() *Funds {
	mapFunds := make(map[string]Fund)

	for _, investiment := range *ii {
		var value float64
		var checkpoint int
		if investiment.CheckpointValue == nil {
			value = investiment.Amount
		} else {
			value = *investiment.CheckpointValue
			checkpoint = 1
		}

		fund, ok := mapFunds[investiment.Cnpj]
		if ok {
			fund.Amount = fund.Amount + investiment.Amount
			fund.Value = fund.Value + value
			fund.Share = fund.Share + 1
			fund.Checkpoints = fund.Checkpoints + checkpoint
			mapFunds[investiment.Cnpj] = fund
		} else {
			mapFunds[investiment.Cnpj] = Fund{
				Name:         investiment.Name,
				Cnpj:         investiment.Cnpj,
				Bank:         investiment.Bank,
				BoxName:      investiment.Box,
				CategoryName: investiment.Category,
				Amount:       investiment.Amount,
				Value:        value,
				Share:        1,
				Checkpoints:  checkpoint,
			}
		}
	}

	funds := Funds{}
	for _, fund := range mapFunds {
		funds = append(funds, fund)
	}

	return &funds
}
