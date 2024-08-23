package controllers

import (
	"log"
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/cmd/web/views/components/trips"
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/routes"
)

type Trips struct {
	queries *database.Queries
	store   auth.UserSessionStore
	logger  logging.Logger
}

func NewTrips(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *Trips {
	return &Trips{
		queries: q,
		store:   u,
		logger:  l,
	}
}

func (c *Trips) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	}
}

func (c *Trips) Get(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		// redirect to trips if already in q params
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}
	user, _ := c.queries.GetUserByEmail(r.Context(), email)
	trips, _ := c.queries.ListUserTrips(r.Context(), user.UserID)
	logger := log.Default()
	logger.Println("total trips: ", len(trips))

	RenderComponent(views.Trips(models.NewTripsModel(email, &trips)), w, r)
}

func (c *Trips) MapNewTrips(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.GetNewTripForm(w, r)
	}
}

func (c *Trips) GetNewTripForm(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		// redirect to trips if already in q params
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}
	RenderComponent(trips.NewTrip(email), w, r)
}
