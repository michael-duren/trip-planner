package controllers

import (
	"net/http"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
)

type Home struct {
	queries *database.Queries
}

func NewHome(q *database.Queries) *Home {
	return &Home{
		queries: q,
	}
}

func (c *Home) Get(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	RenderComponent(views.Home(), w, r)
}
