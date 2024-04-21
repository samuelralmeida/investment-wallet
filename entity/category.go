package entity

type Category struct {
	ID              int
	Name            string
	EstimatedMonths int
	RateIndicated   float64
	Rules           []string
	Notes           []string
	SubCategory     []SubCategory
}

type SubCategory struct {
	ID         int
	Name       string
	CategoryID int
	Rules      []string
	Notes      []string
}

type Categories []Category

func (cs *Categories) FindByID(id int) *Category {
	for _, c := range *cs {
		if c.ID == id {
			return &c
		}
	}
	return nil
}

func (cs *Categories) AddCategory(category Category) {
	*cs = append(*cs, category)
}

func (cs *Categories) AddSubCategory(categoryID int, subCategory SubCategory) {
	for i := range *cs {
		if (*cs)[i].ID == categoryID {
			(*cs)[i].SubCategory = append((*cs)[i].SubCategory, subCategory)
			return
		}
	}
}
