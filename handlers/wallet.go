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

	return renderTemplate(c, http.StatusOK, "wallet.html", wallets)
}

func (h *Handlers) SaveWallet(c echo.Context) error {
	name := c.FormValue("name")
	fmt.Println("Name", name)
	if name == "" {
		return ApiError{StatusCode: http.StatusBadRequest, Err: nil, Msg: "nome da carteira inválido"}
	}

	wallet := &entity.Wallet{Name: name}
	err := h.Services.SaveWallet(c.Request().Context(), wallet)
	if err != nil {
		return ApiError{StatusCode: http.StatusInternalServerError, Err: err, Msg: "erro ao salvar uma carteira"}
	}

	return renderTemplate(c, http.StatusOK, "wallet_li.html", wallet)
}
