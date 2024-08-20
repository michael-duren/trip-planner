package handlers

import "trip-planner/internal/database"

// A handler is specific to api endpoints
// use a controller if dealing with views or components
type Handlers struct {
	Queries *database.Queries
}

func NewHandlers(q *database.Queries) (h *Handlers) {
	return &Handlers{
		Queries: q,
	}
}
