package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samuelralmeida/investiment-calc/entity"
	"github.com/samuelralmeida/investiment-calc/templates"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		investiments := []entity.Investiment{
			{
				ID:       "um",
				Date:     time.Now(),
				Box:      "estabilidade",
				Category: "baunilha",
				Name:     "capitania premium 45",
				Bank:     "btg",
				Cnpj:     "20146294000171",
				Amount:   303750,
				Pocket:   "principal",
			},
			{
				ID:       "dois",
				Date:     time.Now(),
				Box:      "diversificação",
				Category: "viés macro",
				Name:     "kinea atlas",
				Bank:     "btg",
				Cnpj:     "29762315000158",
				Amount:   313200,
				Pocket:   "principal",
			},
		}
		t, _ := template.ParseFS(templates.FS, "investiments.html")
		t.Execute(w, investiments)
	})

	r.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		for key, values := range r.Form {
			fmt.Println("key", key, "value", values[0])
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	log.Println("running in port 3000...")
	http.ListenAndServe(":3000", r)
}
