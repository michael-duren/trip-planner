package controllers

import (
	"trip-planner/internal/database"

	"github.com/gorilla/sessions"
)

type Controllers struct {
	*Home
	*Trips
	*TripBuilder
	*Auth
}

func NewControllers(q *database.Queries, s *sessions.CookieStore) *Controllers {
	return &Controllers{
		NewHome(q, s),
		NewTrips(q, s),
		NewTripBuilder(q, s),
		NewAuth(q, s),
	}
}
