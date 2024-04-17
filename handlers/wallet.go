package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samuelralmeida/investment-wallet/entity"
)

type IWalletService interface {
	ListWallets(ctx context.Context) ([]entity.Wallet, error)
	SaveWallet(ctx context.Context, wallet *entity.Wallet) error
}

func (h *Handlers) RenderListWallets(c echo.Context) error {
	ctx := c.Request().Context()

	wallets, err := h.Services.ListWallets(ctx)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "list wallets err"}
	}

	html := "<div>"
	for _, wallet := range wallets {
		w := fmt.Sprintf("<p>%d - %s</p>", wallet.ID, wallet.Name)
		html = html + w
	}
	html = html + "</div>"

	return c.HTML(200, html)
}

func (h *Handlers) SaveWallet(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return ApiError{StatusCode: http.StatusBadRequest, Err: nil, Msg: "invalid wallet name"}
	}

	wallet := &entity.Wallet{Name: name}
	err := h.Services.SaveWallet(c.Request().Context(), wallet)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "save wallet err"}
	}

	html := "<div>"
	w := fmt.Sprintf("<p>%d - %s</p>", wallet.ID, wallet.Name)
	html = html + w
	html = html + "</div>"

	return c.HTML(http.StatusCreated, html)

}
