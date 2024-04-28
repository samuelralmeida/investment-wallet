package entity

type Fund struct {
	ID        int
	Name      string
	Cnpj      string
	Benchmark string
	Category  Category
	Bank      string
	MinValue  int
	Notes     []string
}
