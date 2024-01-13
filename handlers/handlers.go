package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/samuelralmeida/investiment-calc/entity"
	"github.com/samuelralmeida/investiment-calc/templates"
)

type InvestimentServiceInterface interface {
	ListInvestiments(ctx context.Context) (*[]entity.Investiment, error)
	CreateInvestiment(ctx context.Context, investiment *entity.Investiment) error
	CreateInvestmentCheckpoint(ctx context.Context, investimentCheckpoints *[]entity.InvestimentCheckpoint) error
}

type handlers struct {
	Service InvestimentServiceInterface
}

func New(service InvestimentServiceInterface) *handlers {
	return &handlers{Service: service}
}

func (h *handlers) RenderInvestimentsList(w http.ResponseWriter, r *http.Request) {
	investiments, err := h.Service.ListInvestiments(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "error to fetch investiments", http.StatusInternalServerError)
		return
	}

	t, _ := template.ParseFS(templates.FS, "investiments.html")
	t.Execute(w, investiments)
}

func (h *handlers) SaveInvestimentCheckpoints(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	checkpoints := []entity.InvestimentCheckpoint{}

	for key, values := range r.Form {
		if len(values) != 1 {
			log.Println("each investiment has only one checkpoint")
			http.Error(w, "invalid value", http.StatusBadRequest)
			return
		}

		value, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			log.Println(err)
			http.Error(w, "value must be a number", http.StatusBadRequest)
			return
		}

		checkpoints = append(checkpoints, entity.InvestimentCheckpoint{
			InvestimentID: key,
			Amount:        value,
		})
	}

	err := h.Service.CreateInvestmentCheckpoint(r.Context(), &checkpoints)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to save checkpoints", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/investiments", http.StatusFound)
}
