package controllers

import (
	"net/http"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
)

type HomeController struct {
	queries *database.Queries
}

func NewHomeController(q *database.Queries) *HomeController {
	return &HomeController{
		queries: q,
	}
}

func (c *HomeController) Get(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	RenderComponent(views.Home(), w, r)
}
