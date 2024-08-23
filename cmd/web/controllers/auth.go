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
		c.GetLogin(w, r)
	case "POST":
		c.PostLogin(w, r)
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
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Fetch the user by email
	user, err := c.queries.GetUserByEmail(r.Context(), email)
	if errors.Is(err, sql.ErrNoRows) {
		RenderComponent(authforms.LoginForm(&authforms.LoginValidationErrors{"email": "Email or Password is incorrect"}), w, r)
		return
	}

	// Compare the provided password with the stored hash
	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		RenderComponent(authforms.LoginForm(&authforms.LoginValidationErrors{"password": "Email or Password is incorrect"}), w, r)
		return
	}

	// Create user session
	err = c.store.CreateUserSession(r, w, &user)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	// Redirect after successful login
	w.Header().Set("HX-Redirect", routes.Trips)
	w.WriteHeader(http.StatusOK)
}

func (c *Auth) MapLogout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := c.store.DeleteUserSession(r, w)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Redirect after successful logout
		http.Redirect(w, r, routes.Home, http.StatusSeeOther)
	}
}

func (c *Auth) MapRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.GetRegister(w, r)
	case "POST":
		c.PostRegister(w, r)
	}
}

func (c *Auth) GetRegister(w http.ResponseWriter, r *http.Request) {
	RenderComponent(authforms.RegisterForm(nil), w, r)
}

func (c *Auth) PostRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.logger.Info("unable to parse form from postregister")
		http.Error(w, "Bad Request", http.StatusFound)
	}

	// TODO: find better way to parse forms
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repwd := r.FormValue("re-password")

	formVals := authforms.RegisterFormValues{
		Username:   username,
		Email:      email,
		Password:   password,
		RePassword: repwd,
	}

	// TODO: Find better validation procedure
	if password != repwd {
		RenderComponent(
			authforms.RegisterForm(authforms.NewRegisterFormProps(&authforms.RegisterFormValidationErrors{"Password": "Passwords do not match"}, &formVals)),
			w,
			r,
		)
		return
	}
	_, err = c.queries.GetUserByEmail(r.Context(), email)
	if !errors.Is(err, sql.ErrNoRows) {
		RenderComponent(
			authforms.RegisterForm(authforms.NewRegisterFormProps(&authforms.RegisterFormValidationErrors{"Email": "A user with this email address already exists"}, &formVals)),
			w,
			r,
		)
		return
	}
	_, err = c.queries.GetUserByUsername(r.Context(), username)
	if !errors.Is(err, sql.ErrNoRows) {
		registerValidationErrors := &authforms.RegisterFormValidationErrors{"Username": "A user with this username address already exists"}
		RenderComponent(authforms.RegisterForm(authforms.NewRegisterFormProps(registerValidationErrors, &formVals)), w, r)
		return
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

	err = c.store.CreateUserSession(r, w, &newDbUser)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	// Redirect after successful registration
	w.Header().Set("HX-Redirect", routes.Trips)
	w.WriteHeader(http.StatusCreated)
}
