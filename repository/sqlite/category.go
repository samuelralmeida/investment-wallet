package sqlite

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/investment-wallet/entity"
)

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
			SubCategories:   []entity.SubCategory{subCategory},
		}

		categories = append(categories, *category)
	}

	return categories, nil
}

func (r *Repository) selectCategoryBySubCategoryID(ctx context.Context, subCategoryID int) (*entity.Category, error) {
	query := `
		select
			c.id, c.name, c.estimated_months, c.rate_indicated, c.rules, c.notes,
			sc.id as subcategory_id, sc.name, sc.rules, sc.notes
		from category c
		join sub_category sc on sc.category_id = c.id
		where sc.id = ?
	`

	row := r.db.QueryRowContext(ctx, query, subCategoryID)
	category := entity.Category{}
	err := row.Scan(
		&category.ID, &category.Name, &category.EstimatedMonths, &category.RateIndicated,
		&category.Rules, &category.Notes,
	)

	if err != nil {
		return nil, fmt.Errorf("scan category by subcategory id: %w", err)
	}

	return &category, nil
}
