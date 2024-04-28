package handlers

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/samuelralmeida/investment-wallet/templates"
)

type IService interface {
	IWalletService
	IFundService
}

type Handlers struct {
	Services IService
}

func New(services IService) *Handlers {
	return &Handlers{
		Services: services,
	}
}

func renderTemplate(c echo.Context, code int, templateName string, data interface{}) error {
	t, err := template.ParseFS(templates.FS, templateName)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	return c.HTMLBlob(code, buf.Bytes())
}
