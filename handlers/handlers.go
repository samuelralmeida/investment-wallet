package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

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

func (h *handlers) RenderInvestimentNew(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFS(templates.FS, "investiment.html")
	t.Execute(w, nil)
}

func (h *handlers) SaveInvestiment(w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "amount must be a number", http.StatusBadRequest)
		return
	}

	investiment := &entity.Investiment{}

	investiment.Name = r.FormValue("fund")
	investiment.Cnpj = r.FormValue("cnpj")
	investiment.Box = r.FormValue("box")
	investiment.Category = r.FormValue("category")
	investiment.Bank = r.FormValue("bank")
	investiment.Wallet = r.FormValue("wallet")
	investiment.Date = time.Now()
	investiment.Amount = amount

	err = h.Service.CreateInvestiment(r.Context(), investiment)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to save investiment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/investiments", http.StatusFound)
}
