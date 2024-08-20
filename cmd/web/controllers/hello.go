package controllers

import (
	"log"
	"net/http"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
)

type HelloWorldController struct {
	queries *database.Queries
}

func NewHelloWorldController(q *database.Queries) *HelloWorldController {
	return &HelloWorldController{
		queries: q,
	}
}

func (c *HelloWorldController) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := views.HelloPost(name)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
