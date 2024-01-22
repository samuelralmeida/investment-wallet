package entity

import (
	"time"
)

type Fund struct {
	ID string
	Name string
	Cnpj string
	Box string
	Flavor string
	Bank string
	MinValue float64
	DeletedAt *time.Time
}

type Funds []Fund

/*
// Fund is the collection of shares of the same fund
type Fund struct {
	Name         string
	Cnpj         string
	Bank         string
	BoxName      string
	CategoryName string
	Amount       float64
	Value        float64
	Share        int
	Checkpoints  int
}

type Funds []Fund

// Boxes group funds according to the type of box
func (f *Funds) Boxes() *Boxes {
	mapBoxes := make(map[string]Box)

	for _, fund := range *f {
		box, ok := mapBoxes[fund.BoxName]
		if ok {
			box.Amount = box.Amount + fund.Amount
			box.Value = box.Value + fund.Value
			box.QuantityFunds = box.QuantityFunds + 1
			box.QuantityCheckpoints = box.QuantityCheckpoints + fund.Checkpoints
			box.QuantityShares = box.QuantityShares + fund.Share
			mapBoxes[fund.BoxName] = box
		} else {
			mapBoxes[fund.BoxName] = Box{
				Name:                fund.BoxName,
				Amount:              fund.Amount,
				Value:               fund.Value,
				QuantityFunds:       1,
				QuantityCheckpoints: fund.Checkpoints,
				QuantityShares:      fund.Share,
			}
		}
	}

	boxes := Boxes{}
	for _, box := range mapBoxes {
		boxes = append(boxes, box)
	}

	return &boxes
}

// Boxes group funds according to the type of category
func (f *Funds) Categories() *Categories {
	mapCategories := make(map[string]Category)

	for _, fund := range *f {
		key := fmt.Sprintf("%s:%s", fund.BoxName, fund.CategoryName)
		category, ok := mapCategories[key]
		if ok {
			category.Amount = category.Amount + fund.Amount
			category.Value = category.Value + fund.Value
			category.QuantityFunds = category.QuantityFunds + 1
			category.QuantityCheckpoints = category.QuantityCheckpoints + fund.Checkpoints
			category.QuantityShares = category.QuantityShares + fund.Share
			mapCategories[key] = category
		} else {
			mapCategories[key] = Category{
				Name:                fund.CategoryName,
				Amount:              fund.Amount,
				Value:               fund.Value,
				QuantityFunds:       1,
				Box:                 fund.BoxName,
				QuantityCheckpoints: fund.Checkpoints,
				QuantityShares:      fund.Share,
			}
		}
	}

	categories := Categories{}
	for _, category := range mapCategories {
		categories = append(categories, category)
	}

	return &categories
}

func (f *Funds) TotalAmount() float64 {
	var total float64
	for _, fund := range *f {
		total = total + fund.Amount
	}
	return total
}

func (f *Funds) TotalValue() float64 {
	var total float64
	for _, fund := range *f {
		total = total + fund.Value
	}
	return total
}

func (f *Funds) TotalShares() int {
	var total int
	for _, fund := range *f {
		total = total + fund.Share
	}
	return total
}

func (f *Funds) TotalCheckpoints() int {
	var total int
	for _, fund := range *f {
		total = total + fund.Checkpoints
	}
	return total
}
*/