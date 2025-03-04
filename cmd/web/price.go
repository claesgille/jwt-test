package web

import (
	"jwt-test/internal/model"
	"log"
	"net/http"
)

func AllPriceHandler(w http.ResponseWriter, r *http.Request) {
	component := AllPrices(model.Draws)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
func PriceHandler(id string, w http.ResponseWriter, r *http.Request) {
	draw := getDraw(id)
	component := DrawTemplate(draw)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
