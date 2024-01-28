package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/samuelralmeida/investment-wallet/database"
	"github.com/samuelralmeida/investment-wallet/handlers"
	"github.com/samuelralmeida/investment-wallet/repository"
	"github.com/samuelralmeida/investment-wallet/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db, err := database.NewSqliteConnection()
	if err != nil {
		panic(err)
	}

	repository := repository.New(db)
	service := service.New(repository)
	handlers := handlers.New(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/funds/new", handlers.RenderNewFund)
	r.Post("/funds/new", handlers.NewFund)

	r.Get("/investments/new", handlers.RenderNewInvestment)
	r.Post("/investments/new", handlers.NewInvestment)

	r.Get("/checkpoints/new", handlers.RenderNewCheckpoint)
	r.Post("/checkpoints/new", handlers.NewCheckpoint)

	r.Get("/wallet/{name}", handlers.Wallet)
	r.Get("/calculate/{name}", handlers.Calculate)

	log.Println("running in port 3000...")
	http.ListenAndServe(":3000", r)
}
