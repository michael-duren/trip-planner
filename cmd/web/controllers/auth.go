package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"trip-planner/cmd/web/views/components/authforms"
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/routes"
)

type Auth struct {
	queries *database.Queries
	store   auth.UserSessionStore
	logger  logging.Logger
}

func NewAuth(q *database.Queries, u auth.UserSessionStore, l logging.Logger) *Auth {
	return &Auth{
		queries: q,
		store:   u,
		logger:  l,
	}
}

func (c *Auth) MapLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	}
}

func (c *Auth) GetLogin(w http.ResponseWriter, r *http.Request) {
	RenderComponent(authforms.LoginForm(nil), w, r)
}

func (c *Auth) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.logger.Info("unable to parse form from postregister")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	pwdHash, err := auth.HashPassword(password)
	if err != nil {
		c.logger.Panicf("unable to hash pwd in PostLogin method. Ending program, error: %s", err.Error())
	}
	user, err := c.queries.LoginUser(r.Context(), database.LoginUserParams{
		Email:    email,
		Password: pwdHash,
	})
	if errors.Is(err, sql.ErrNoRows) {
		RenderComponent(authforms.LoginForm(&authforms.LoginValidationErrors{"email": "email or password is incorrect"}), w, r)
		return
	}
	c.store.CreateUserSession(r, w, &user)
	http.Redirect(w, r, routes.Trips, http.StatusOK)
}

func (c *Auth) MapRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	}
}

func (c *Auth) GetRegister(w http.ResponseWriter, r *http.Request) {
	RenderComponent(authforms.RegisterForm(nil), w, r)
}

func (c *Auth) PostRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.logger.Info("unable to parse form from postregister")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	// TODO: find better way to parse forms
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repwd := r.FormValue("re-password")

	// TODO: Find better validation procedure
	if password != repwd {
		RenderComponent(authforms.RegisterForm(&authforms.RegisterFormValidationErrors{"Password": "Passwords do not match"}), w, r)
	}
	_, err = c.queries.GetUserByEmail(r.Context(), email)
	if !errors.Is(err, sql.ErrNoRows) {
		RenderComponent(authforms.RegisterForm(&authforms.RegisterFormValidationErrors{"Email": "A user with this email address already exists"}), w, r)
	}
	_, err = c.queries.GetUserByUsername(r.Context(), username)
	if !errors.Is(err, sql.ErrNoRows) {
		RenderComponent(authforms.RegisterForm(&authforms.RegisterFormValidationErrors{"Username": "A user with this username address already exists"}), w, r)
	}

	// create user
	pwdHash, err := auth.HashPassword(password)
	if err != nil {
		c.logger.Warnf("Error in PostRegister hashing password", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	newDbUser, err := c.queries.RegisterUser(r.Context(), database.RegisterUserParams{
		Email:    email,
		Username: username,
		Password: pwdHash,
	})
	if err != nil {
		c.logger.Warnf("Error in PostRegister creating user", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	c.store.CreateUserSession(r, w, &newDbUser)

	http.Redirect(w, r, routes.Trips, http.StatusCreated)
}
