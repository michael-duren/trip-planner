package controllers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func RenderComponent(component templ.Component, w http.ResponseWriter, r *http.Request) {
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in controller: %e", err)
	}
}
