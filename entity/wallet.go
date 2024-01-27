package entity

type FundDetail struct {
	Fund        Fund
	Investments Investments
	Checkpoint  Checkpoint
}

type Wallet struct {
	Name        string
	FundsDetail []FundDetail
}

func (fd *FundDetail) TotalInvestedValue() float64 {
	var total float64
	for _, investment := range fd.Investments {
		total = total + investment.Value
	}
	return total
}

func (w *Wallet) TotalInvestedValue() float64 {
	var total float64
	for _, fund := range w.FundsDetail {
		total = total + fund.TotalInvestedValue()
	}
	return total
}

func (w *Wallet) TotalCurrentValue() float64 {
	var total float64
	for _, fund := range w.FundsDetail {
		total = total + fund.Checkpoint.Value
	}
	return total
}

func (w *Wallet) RatioTotalValue() float64 {
	return w.TotalCurrentValue() / w.TotalInvestedValue()
}

func (w *Wallet) Boxes() Boxes {
	boxesMap := make(map[string]Box)
	for _, fund := range w.FundsDetail {
		boxName := fund.Fund.Box
		box := boxesMap[boxName]
		boxesMap[boxName] = Box{
			Name:          boxName,
			InvestedValue: box.InvestedValue + fund.TotalInvestedValue(),
			CurrentValue:  box.CurrentValue + fund.Checkpoint.Value,
			FundsDetail:   append(box.FundsDetail, fund),
		}
	}

	boxes := Boxes{}
	for _, box := range boxesMap {
		boxes = append(boxes, box)
	}

	return boxes
}
