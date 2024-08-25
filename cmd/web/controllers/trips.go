package controllers

import (
	"database/sql"
	"errors"
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
	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	trips, _ := c.queries.ListUserTrips(r.Context(), user.UserID)
	logger := log.Default()
	logger.Println("total trips: ", len(trips))

	RenderComponent(views.Trips(models.NewTripsModel(user, &trips)), w, r)
}

func (c *Trips) MapNewTrips(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.GetNewTripForm(w, r)
	case "POST":
		c.PostNewTripForm(w, r)
	}
}

func (c *Trips) GetNewTripForm(w http.ResponseWriter, r *http.Request) {
	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}

	RenderComponent(trips.NewTrip(trips.NewNewTripProps("")), w, r)
}

func (c *Trips) PostNewTripForm(w http.ResponseWriter, r *http.Request) {
	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		http.Redirect(w, r, routes.Home, http.StatusTemporaryRedirect)
		return
	}
	c.logger.Info(user.ToString())
	c.logger.Info("In PostNewTripForm")
	err := r.ParseForm()
	if err != nil {
		c.logger.Info("unable to parse form from postregister")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tripName := r.FormValue("trip-name")
	_, err = c.queries.GetTripByName(r.Context(), tripName)
	if !errors.Is(err, sql.ErrNoRows) {
		c.logger.Info("Trip with that name already exists")
		RenderComponent(trips.NewTrip(trips.NewNewTripProps("Trip with a name already exists")), w, r)
		return
	}

	newTrip, err := c.queries.CreateTrip(r.Context(), database.CreateTripParams{
		UserID: user.UserID,
		Name:   tripName,
	})
	if err != nil {
		http.Redirect(w, r, routes.Home, http.StatusInternalServerError)
		return
	}

	redirectUrl := routes.QueryParamBuilder(routes.TripBuilder, routes.QueryParams{"trip-id": string(newTrip.TripID)})
	// Redirect after successful registration
	w.Header().Set("HX-Redirect", redirectUrl)
	w.WriteHeader(http.StatusCreated)
}
