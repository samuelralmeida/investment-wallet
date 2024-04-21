package sqlite

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/samuelralmeida/investment-wallet/entity"
)

type jsonField []string

func (jf *jsonField) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	data, ok := value.(string)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	err := json.Unmarshal([]byte(data), &jf)
	if err != nil {
		return fmt.Errorf("scan json field: %w", err)
	}

	return nil
}

func (jf jsonField) Value() (driver.Value, error) {
	b := new(strings.Builder)
	err := json.NewEncoder(b).Encode(jf)
	return b.String(), err
}

type sqliteCategory struct {
	ID               int
	Name             string
	EstimatedMonths  int
	RateIndicated    float64
	Rules            jsonField
	Notes            jsonField
	SubCategoryID    int
	SubCategoryName  string
	SubCategoryRules jsonField
	SubCategoryNotes jsonField
}

func (r *Repository) SelectCategories(ctx context.Context) ([]entity.Category, error) {
	query := `
		select
			c.id, c.name, c.estimated_months, c.rate_indicated, c.rules, c.notes,
			sc.id as subcategory_id, sc.name, sc.rules, sc.notes
		from category c
		join sub_category sc on sc.category_id = c.id
		order by c.id asc
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("select categories: %w", err)
	}

	categories := entity.Categories{}
	for rows.Next() {
		var sc sqliteCategory

		err := rows.Scan(
			&sc.ID, &sc.Name, &sc.EstimatedMonths, &sc.RateIndicated, &sc.Rules, &sc.Notes,
			&sc.SubCategoryID, &sc.SubCategoryName, &sc.SubCategoryRules, &sc.SubCategoryNotes,
		)

		if err != nil {
			return nil, fmt.Errorf("scan category row: %w", err)
		}

		subCategory := entity.SubCategory{
			ID:         sc.SubCategoryID,
			Name:       sc.SubCategoryName,
			CategoryID: sc.ID,
			Rules:      sc.SubCategoryRules,
			Notes:      sc.SubCategoryNotes,
		}

		category := categories.FindByID(sc.ID)
		if category != nil {
			categories.AddSubCategory(sc.ID, subCategory)
			continue
		}

		category = &entity.Category{
			ID:              sc.ID,
			Name:            sc.Name,
			EstimatedMonths: sc.EstimatedMonths,
			RateIndicated:   sc.RateIndicated,
			Rules:           sc.Rules,
			Notes:           sc.Notes,
			SubCategory:     []entity.SubCategory{subCategory},
		}

		categories = append(categories, *category)
	}

	fmt.Println(categories)

	return categories, nil

}
