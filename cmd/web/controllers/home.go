package controllers

import (
	"net/http"
	"trip-planner/cmd/web/views"

	"github.com/a-h/templ"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	templ.Handler(views.Home())
}
