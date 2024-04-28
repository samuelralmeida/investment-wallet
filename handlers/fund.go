package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelralmeida/investment-wallet/entity"
)

type IFundService interface {
	ListCategories(ctx context.Context) ([]entity.Category, error)
	ListFunds(ctx context.Context) ([]entity.Fund, error)
	SaveFund(ctx context.Context, fund *entity.Fund, subCategoryID int) error
}

func (h *Handlers) RenderListFunds(c echo.Context) error {
	ctx := c.Request().Context()

	categories, err := h.Services.ListCategories(ctx)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "list categories err"}
	}

	funds, err := h.Services.ListFunds(ctx)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "list funds err"}
	}

	data := struct {
		Categories []entity.Category
		Funds      []entity.Fund
	}{
		Categories: categories,
		Funds:      funds,
	}

	return renderTemplate(c, http.StatusOK, "fund.html", data)
}

func (h *Handlers) SaveFund(c echo.Context) error {
	type input struct {
		Name          string   `form:"name"`
		Cnpj          string   `form:"cnpj"`
		Bank          string   `form:"bank"`
		Benchmark     string   `form:"benchmark"`
		Notes         []string `form:"notes"`
		MinValue      float64  `form:"minValue"`
		SubCategoryID int      `form:"subCategory"`
	}

	payload := new(input)
	err := c.Bind(payload)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "erro para validar campos"}
	}

	fund := &entity.Fund{
		Name: payload.Name, Cnpj: payload.Cnpj, Benchmark: payload.Benchmark,
		Bank: payload.Bank, Notes: payload.Notes, MinValue: int(payload.MinValue * 100)}
	err = h.Services.SaveFund(c.Request().Context(), fund, payload.SubCategoryID)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "erro ao salvar um fundo"}
	}

	return c.Redirect(http.StatusSeeOther, "/funds")
}

func (h *Handlers) ElementFundNote(c echo.Context) error {
	return renderTemplate(c, http.StatusOK, "fund_note.html", nil)
}
