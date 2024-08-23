package controllers

import (
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
)

type Controllers struct {
	*Home
	*Trips
	*TripBuilder
	*Auth
}

func NewControllers(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *Controllers {
	return &Controllers{
		NewHome(q, u, l),
		NewTrips(q, u, l),
		NewTripBuilder(q, u, l),
		NewAuth(q, u, l),
	}
}
