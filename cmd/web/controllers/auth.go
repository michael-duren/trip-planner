package controllers

import (
	"net/http"
	"trip-planner/cmd/web/views/components/authforms"
	"trip-planner/internal/database"

	"github.com/gorilla/sessions"
)

type Auth struct {
	queries *database.Queries
	store   *sessions.CookieStore
}

func NewAuth(q *database.Queries, s *sessions.CookieStore) *Auth {
	return &Auth{
		queries: q,
		store:   s,
	}
}

func (c *Auth) MapLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.GetLogin(w, r)
	}
}

func (c *Auth) GetLogin(w http.ResponseWriter, r *http.Request) {
	RenderComponent(authforms.LoginForm(), w, r)
}
