package entity

type Box struct {
	Name          string
	InvestedValue float64
	CurrentValue  float64
	FundsDetail   []FundDetail
}

type Boxes []Box
