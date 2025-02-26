package web

import (
	"log"
	"net/http"
	"strconv"
)

type Draw struct {
	Id     string
	Status string
	Name   string
}

var draws = []Draw{
	{Id: "1", Status: "Closed", Name: "Dragning 2022"},
	{Id: "2", Status: "Closed", Name: "Dragning 2023"},
	{Id: "3", Status: "Open", Name: "Dragning 2024"},
}

func AllDrawsHandler(w http.ResponseWriter, r *http.Request) {
	component := AllDraws(draws)
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
	draws = append(draws, Draw{Id: strconv.Itoa(len(draws) + 1), Status: "New", Name: name})
}

func getDraw(id string) *Draw {
	for _, draw := range draws {
		if draw.Id == id {
			return &draw
		}
	}
	return nil
}
