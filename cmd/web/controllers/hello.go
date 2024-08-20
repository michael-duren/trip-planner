package controllers

import (
	"log"
	"net/http"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
)

type HelloWorld struct {
	queries *database.Queries
}

func NewHelloWorld(q *database.Queries) *HelloWorld {
	return &HelloWorld{
		queries: q,
	}
}

func (c *HelloWorld) Post(w http.ResponseWriter, r *http.Request) {
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
