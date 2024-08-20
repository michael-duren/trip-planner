package controllers

import (
	"net/http"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
	"trip-planner/internal/server/routes"
)

type Trips struct {
	queries *database.Queries
}

func NewTrips(q *database.Queries) *Trips {
	return &Trips{
		queries: q,
	}
}

func (c *Trips) Get(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		// redirect to trips if already in q params
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}

	RenderComponent(views.Trips(email), w, r)
}
