package server

import (
	"encoding/json"
	"log"
	"net/http"
	"trip-planner/cmd/web"
	"trip-planner/cmd/web/controllers"
	"trip-planner/cmd/web/views"
	"trip-planner/internal/auth"
	"trip-planner/internal/server/handlers"
	"trip-planner/internal/server/routes"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	h := handlers.NewHandlers(s.db.UseQueries())
	store := auth.CreateStore()
	c := controllers.NewControllers(s.db.UseQueries(), store)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	// api
	mux.HandleFunc("/test-endpoint", h.HelloWorldHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc(routes.Login, c.Auth.MapLogin)

	// pages
	// home
	mux.HandleFunc(routes.Home, c.Home.Map)
	// trips
	mux.HandleFunc(routes.Trips, c.Trips.Map)
	mux.HandleFunc(routes.NewTrips, c.Trips.MapNewTrips)
	// tripbuilder
	mux.HandleFunc(routes.TripBuilder, c.TripBuilder.Map)

	// examples
	mux.Handle("/web", templ.Handler(views.HelloForm()))

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
