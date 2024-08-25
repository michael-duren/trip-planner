package controllers

import (
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/routes"
)

type TripBuilder struct {
	queries *database.Queries
	store   auth.UserSessionStore
	logger  logging.Logger
}

func NewTripBuilder(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *TripBuilder {
	return &TripBuilder{
		queries: q,
		store:   u,
		logger:  l,
	}
}

func (c *TripBuilder) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	}
}

func (c *TripBuilder) Get(w http.ResponseWriter, r *http.Request) {
	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}

	// TODO: rm hardcoded values
	RenderComponent(views.TripBuilder(&models.TripBuildersModel{MainLayoutModel: models.NewMainLayout(user.Email), Tripname: "trip", Tripid: 1}), w, r)
}
