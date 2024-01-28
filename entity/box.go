package entity

type Box struct {
	Name          string
	InvestedValue float64
	CurrentValue  float64
	FundsDetail   []FundDetail
}

func (b *Box) Income() float64 {
	return ((b.CurrentValue / b.InvestedValue) * 100) - 100
}

type Boxes []Box
