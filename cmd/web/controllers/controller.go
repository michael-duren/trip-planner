package controllers

import "trip-planner/internal/database"

type Controllers struct {
	*HelloWorld
	*Home
	*Trips
}

func NewControllers(q *database.Queries) *Controllers {
	return &Controllers{
		NewHelloWorld(q),
		NewHome(q),
		NewTrips(q),
	}
}
