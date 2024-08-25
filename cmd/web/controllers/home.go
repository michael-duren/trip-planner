package controllers

import (
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/routes"
)

type Home struct {
	queries *database.Queries
	store   auth.UserSessionStore
	logger  logging.Logger
}

func NewHome(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *Home {
	return &Home{
		queries: q,
		store:   u,
		logger:  l,
	}
}

func (c *Home) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	}
}

func (c *Home) Get(w http.ResponseWriter, r *http.Request) {
	user, _ := c.store.GetUserFromSession(r, w)
	if user == nil {
		RenderComponent(views.Home(*models.NewHomeModel(nil)), w, r)
		return
	}
	c.logger.Info(user.ToString())
	http.Redirect(w, r, routes.Trips, http.StatusTemporaryRedirect)
}
