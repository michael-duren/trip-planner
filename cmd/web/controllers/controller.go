package controllers

import "trip-planner/internal/database"

type Controllers struct {
	*HelloWorldController
	*HomeController
}

func NewControllers(q *database.Queries) *Controllers {
	return &Controllers{
		NewHelloWorldController(q),
		NewHomeController(q),
	}
}
