package entity

type Category struct {
	Name                string
	Box                 string
	Amount              float64
	Value               float64
	QuantityFunds       int
	QuantityCheckpoints int
	QuantityShares      int
}

type Categories []Category
