package controllers

import (
	"log"
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
	"trip-planner/internal/server/routes"

	"github.com/gorilla/sessions"
)

type TripBuilder struct {
	queries *database.Queries
	store   *sessions.CookieStore
}

func NewTripBuilder(q *database.Queries, s *sessions.CookieStore) *TripBuilder {
	return &TripBuilder{
		queries: q,
		store:   s,
	}
}

func (c *TripBuilder) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	}
}

func (c *TripBuilder) Get(w http.ResponseWriter, r *http.Request) {
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

	// TODO: rm hardcoded values
	RenderComponent(views.TripBuilder(&models.TripBuildersModel{MainLayoutModel: models.NewMainLayout(email), Tripname: "trip", Tripid: 1}), w, r)
}
