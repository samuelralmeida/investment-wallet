package handlers

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/samuelralmeida/investiment-calc/entity"
	"github.com/samuelralmeida/investiment-calc/internal/box"
	"github.com/samuelralmeida/investiment-calc/templates"
)

type InvestimentServiceInterface interface {
	ListInvestiments(ctx context.Context) (*[]entity.Investiment, error)
	CreateInvestmentCheckpoint(ctx context.Context, investimentCheckpoints *[]entity.Checkpoint) error
	Calculate(ctx context.Context) (*entity.Wallet, error)
	CreateFund(ctx context.Context, fund *entity.Fund) error
	ListFunds(ctx context.Context) (*entity.Funds, error)
	CreateInvestiment(ctx context.Context, investiment *entity.Investment) error
	CreateCheckpoint(ctx context.Context, checkpoint *entity.Checkpoint2) error
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

	checkpoints := []entity.Checkpoint{}

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

		checkpoints = append(checkpoints, entity.Checkpoint{
			InvestimentID: key,
			Value:         value,
			Date:          time.Now(),
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

// func (h *handlers) SaveInvestiment(w http.ResponseWriter, r *http.Request) {
// 	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "amount must be a number", http.StatusBadRequest)
// 		return
// 	}

// 	investiment := &entity.Investiment{}

// 	investiment.Name = r.FormValue("fund")
// 	investiment.Cnpj = r.FormValue("cnpj")
// 	investiment.Box = r.FormValue("box")
// 	investiment.Category = r.FormValue("category")
// 	investiment.Bank = r.FormValue("bank")
// 	investiment.Wallet = r.FormValue("wallet")
// 	investiment.Amount = amount

// 	err = h.Service.CreateInvestiment(r.Context(), investiment)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "error to save investiment", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/investiments/new", http.StatusFound)
// }

func (h *handlers) Calculate(w http.ResponseWriter, r *http.Request) {
	wallet, err := h.Service.Calculate(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "calcuate", http.StatusInternalServerError)
		return
	}

	t, _ := template.New("calculate.html").Funcs(
		template.FuncMap{
			"money": func(input float64) string {
				return strings.Replace(fmt.Sprintf("%.2f", input), ".", ",", 1)
			},
			"ratio": func(part, total float64) float64 {
				return (part / total) * 100
			},
		},
	).ParseFS(templates.FS, "calculate.html")
	t.Execute(w, wallet)
}

func (h *handlers) RenderNewFund(w http.ResponseWriter, r *http.Request) {
	data := struct{ BoxOptions []string }{BoxOptions: box.OptionsList()}

	t, _ := template.ParseFS(templates.FS, "new-fund.html")
	t.Execute(w, data)
}

func (h *handlers) NewFund(w http.ResponseWriter, r *http.Request) {
	minValue, err := strconv.ParseFloat(r.FormValue("fund-min-value"), 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "min value must be a number", http.StatusBadRequest)
		return
	}

	boxOption := r.FormValue("fund-box")
	isValid, box, flavor := box.ValidateOption(boxOption)
	if !isValid {
		log.Println("box option ivalid:", boxOption)
		http.Error(w, "box option invalid", http.StatusBadRequest)
		return
	}

	fund := &entity.Fund{}

	fund.Name = r.FormValue("fund-name")
	// TODO: validate cnpj
	fund.Cnpj = r.FormValue("fund-cnpj")
	fund.Bank = r.FormValue("fund-bank")
	fund.Box = box
	fund.Flavor = flavor
	fund.MinValue = minValue

	err = h.Service.CreateFund(r.Context(), fund)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to create fund", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/funds/new", http.StatusFound)
}

func (h *handlers) RenderNewInvestment(w http.ResponseWriter, r *http.Request) {
	funds, err := h.Service.ListFunds(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "error to list funds", http.StatusInternalServerError)
		return
	}
	data := struct{ FundOptions *entity.Funds }{FundOptions: funds}

	t, _ := template.ParseFS(templates.FS, "new-investment.html")
	t.Execute(w, data)
}

func (h *handlers) NewInvestment(w http.ResponseWriter, r *http.Request) {
	value, err := strconv.ParseFloat(r.FormValue("investment-value"), 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "investment value must be a number", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.DateOnly, r.FormValue("investment-date"))
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	investment := &entity.Investment{}

	investment.FundID = r.FormValue("investment-fund")
	investment.Wallet = r.FormValue("investment-wallet")
	investment.Date = date
	investment.Value = value

	err = h.Service.CreateInvestiment(r.Context(), investment)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to create investment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/investments/new", http.StatusFound)
}

func (h *handlers) RenderNewCheckpoint(w http.ResponseWriter, r *http.Request) {
	funds, err := h.Service.ListFunds(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "error to list funds", http.StatusInternalServerError)
		return
	}
	data := struct{ FundOptions *entity.Funds }{FundOptions: funds}

	t, _ := template.ParseFS(templates.FS, "new-checkpoint.html")
	t.Execute(w, data)
}

func (h *handlers) NewCheckpoint(w http.ResponseWriter, r *http.Request) {
	value, err := strconv.ParseFloat(r.FormValue("checkpoint-value"), 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "checkpoint value must be a number", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.DateOnly, r.FormValue("checkpoint-date"))
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	checkpoint := &entity.Checkpoint2{}

	checkpoint.FundID = r.FormValue("checkpoint-fund")
	checkpoint.Wallet = r.FormValue("checkpoint-wallet")
	checkpoint.Date = date
	checkpoint.Value = value

	err = h.Service.CreateCheckpoint(r.Context(), checkpoint)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to create checkpoint", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/checkpoints/new", http.StatusFound)
}
