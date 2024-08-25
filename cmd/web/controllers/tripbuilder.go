package controllers

import (
	"net/http"
	"strconv"
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
    tripId := r.URL.Query().Get("trip-id")
    tripIdInt, err := strconv.Atoi(tripId)
    if tripId == "" || err != nil {
        http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
    }
    trip,_ := c.queries.GetTripById(r.Context(), int32(tripIdInt))

	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}

    model := models.NewTripBuildersModel(user, trip)
	RenderComponent(views.TripBuilder(model), w, r)
}
