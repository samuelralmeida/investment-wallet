package entity

type Box struct {
	Name                string
	Amount              float64
	Value               float64
	QuantityFunds       int
	QuantityCheckpoints int
	QuantityShares      int
}

type Boxes []Box

func (bb *Boxes) TotalAmount(boxName string) float64 {
	for _, box := range *bb {
		if box.Name == boxName {
			return box.Amount
		}
	}
	return 0
}

func (bb *Boxes) TotalValue(boxName string) float64 {
	for _, box := range *bb {
		if box.Name == boxName {
			return box.Value
		}
	}
	return 0
}

func (b *Box) RatioValue(totaValue float64) float64 {
	return (b.Value / totaValue) * 100
}

func (b *Box) RatioAmount(totalAmount float64) float64 {
	return (b.Amount / totalAmount) * 100
}
