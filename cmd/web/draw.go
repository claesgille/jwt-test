package web

import (
	"jwt-test/internal/model"
	"log"
	"net/http"
	"strconv"
)

func AllDrawsHandler(w http.ResponseWriter, r *http.Request) {
	component := AllDraws(model.Draws)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
func DrawHandler(id string, w http.ResponseWriter, r *http.Request) {
	draw := getDraw(id)
	component := DrawTemplate(draw)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

func NewDrawsHandler(w http.ResponseWriter, r *http.Request) {
	component := NewDraw()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

func NewDrawFormsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	model.Draws = append(model.Draws, model.Draw{Id: strconv.Itoa(len(model.Draws) + 1), Status: "New", Name: name})
}

func getDraw(id string) *model.Draw {
	for _, draw := range model.Draws {
		if draw.Id == id {
			return &draw
		}
	}
	return nil
}
