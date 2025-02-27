package web

import (
	"log"
	"net/http"
	"strconv"
)

type Price struct {
	Id          string
	Name        string
	Description string
	MainImage   string
}

type Draw struct {
	Id     string
	Status string
	Name   string
	Prices []Price
}

var draws = []Draw{
	{Id: "1", Status: "Closed", Name: "Dragning 2022",
		Prices: []Price{
			{Id: "1", Name: "Röd Vas", Description: "En Tjusig röd vas i äkta hammarplast.", MainImage: "/assets/vase.jpg"},
			{Id: "2", Name: "En räv på isen", Description: "Klassisk naturstudie.", MainImage: "/assets/fox.jpg"},
		},
	},
	{Id: "2", Status: "Closed", Name: "Dragning 2023",
		Prices: []Price{
			{Id: "3", Name: "Klassisk målning", Description: "Skulle kunna vara målad av en av de gamla mästarna.", MainImage: "/assets/painting1.jpg"},
			{Id: "4", Name: "Skallepär", Description: "En bissart dekorerad skalle i sydamerikansk stil.", MainImage: "/assets/skull.jpg"},
		},
	},
	{Id: "3", Status: "Open", Name: "Dragning 2024",
		Prices: []Price{
			{Id: "5", Name: "Persiphone", Description: "En klassisk skulptur. För tankarna till de gamla grekerna.", MainImage: "/assets/sculpture1.jpg"},
			{Id: "6", Name: "Kvinna i färg", Description: "Färgsprakane fotokonst som skulle pryda vilken vägg som helst.", MainImage: "/assets/girl.jpg"},
		},
	},
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
