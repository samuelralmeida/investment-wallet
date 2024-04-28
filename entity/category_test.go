package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategories_AddCategory(t *testing.T) {
	categories := Categories{}
	assert.Len(t, categories, 0)
	categories.AddCategory(Category{ID: 1})
	assert.Len(t, categories, 1)
}

func TestCategories_AddSubCategory(t *testing.T) {
	categories := Categories{Category{ID: 1}}
	assert.Len(t, categories[0].SubCategories, 0)
	categories.AddSubCategory(1, SubCategory{ID: 11})
	assert.Len(t, categories[0].SubCategories, 1)
}
