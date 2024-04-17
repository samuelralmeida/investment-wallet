package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/samuelralmeida/investment-wallet/entity"
	"github.com/samuelralmeida/investment-wallet/handlers"
	"github.com/samuelralmeida/investment-wallet/repository/sqlite"
	"github.com/samuelralmeida/investment-wallet/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := loadConf()

	repository, err := sqlite.NewConnection(conf)
	if err != nil {
		log.Fatal(err)
	}

	s := services.New(repository)
	h := handlers.New(s)

	e := echo.New()
	e.Static("/static", "static")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handlers.HandleApiError())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/wallets", h.RenderListWallets)
	e.POST("/wallets", h.SaveWallet)
	e.Logger.Fatal(e.Start(":1234"))
}

func loadConf() *entity.Conf {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return &entity.Conf{
		SQLITE: entity.Sqlite{
			FILENAME: os.Getenv("SQLITE_FILENAME"),
		},
	}
}
