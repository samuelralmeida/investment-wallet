package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

type ApiError struct {
	StatusCode int
	Err        error
	Msg        string
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

func HandleApiError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (returnErr error) {
			err := next(c)
			if err == nil {
				return nil
			}

			if e, ok := err.(ApiError); ok {
				slog.Error(e.Msg, "error", e.Err.Error())
				return echo.NewHTTPError(e.StatusCode, e.Msg)
			}

			slog.Error("unknown error", "error", err.Error())
			return err
		}
	}
}
