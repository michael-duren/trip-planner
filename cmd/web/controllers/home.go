package controllers

import (
	"fmt"
	"net/http"
	"trip-planner/cmd/web/models"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/database"
	"trip-planner/internal/server/routes"

	"github.com/gorilla/sessions"
)

type Home struct {
	queries *database.Queries
	store   *sessions.CookieStore
}

func NewHome(q *database.Queries, s *sessions.CookieStore) *Home {
	return &Home{
		queries: q,
		store:   s,
	}
}

func (c *Home) Map(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	case "POST":
		c.Post(w, r)
	}
}

func (c *Home) Get(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email != "" {
		// redirect to trips if already in q params
		http.Redirect(w, r, fmt.Sprintf("%s?email=%s", routes.Trips, email), http.StatusFound)
		return
	}

	RenderComponent(views.Home(*models.NewHomeModel("")), w, r)
}

func (c *Home) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.PostForm.Get("email")
	http.Redirect(w, r, fmt.Sprintf("%s?email=%s", routes.Trips, email), http.StatusFound)
}
