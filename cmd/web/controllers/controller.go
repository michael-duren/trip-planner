package controllers

import "trip-planner/internal/database"

type Controllers struct {
	*HelloWorld
	*Home
}

func NewControllers(q *database.Queries) *Controllers {
	return &Controllers{
		NewHelloWorld(q),
		NewHome(q),
	}
}
