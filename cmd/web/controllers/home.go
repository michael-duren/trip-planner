package controllers

import (
	"fmt"
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/routes"
)

type Home struct {
	queries *database.Queries
	store   auth.UserSessionStore
	logger  logging.Logger
}

func NewHome(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *Home {
	return &Home{
		queries: q,
		store:   u,
		logger:  l,
	}
}

func (c *Home) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	}
}

func (c *Home) Get(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email != "" {
		// redirect to trips if already in q params
		http.Redirect(w, r, fmt.Sprintf("%s?email=%s", routes.Trips, email), http.StatusFound)
		return
	}

	RenderComponent(views.Home(*models.NewHomeModel("")), w, r)
}
