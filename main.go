package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samuelralmeida/investiment-calc/database"
	"github.com/samuelralmeida/investiment-calc/handlers"
	"github.com/samuelralmeida/investiment-calc/repository"
	"github.com/samuelralmeida/investiment-calc/service"
)

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

	r.Get("/investiments", handlers.RenderInvestimentsList)
	r.Post("/investiments/checkpoint", handlers.SaveInvestimentCheckpoints)
	r.Get("/investiments/new", handlers.RenderInvestimentNew)
	r.Post("/investiments/new", handlers.SaveInvestiment)

	r.Get("/calculate", handlers.Calculate)

	r.Get("/funds/new", handlers.RenderNewFund)

	log.Println("running in port 3000...")
	http.ListenAndServe(":3000", r)
}
