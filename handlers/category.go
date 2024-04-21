package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelralmeida/investment-wallet/entity"
)

type ICategoryService interface {
	ListCategories(ctx context.Context) ([]entity.Category, error)
}

func (h *Handlers) RenderListCategories(c echo.Context) error {
	ctx := c.Request().Context()

	categories, err := h.Services.ListCategories(ctx)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "list categories err"}
	}

	return renderTemplate(c, http.StatusOK, "category.html", categories)
}
