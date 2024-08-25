package server

import (
	"encoding/json"
	"log"
	"net/http"
	"trip-planner/cmd/web"
	"trip-planner/cmd/web/controllers"
	"trip-planner/internal/auth"
	"trip-planner/internal/logging"
	"trip-planner/internal/server/handlers"
	"trip-planner/internal/server/routes"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	h := handlers.NewHandlers(s.db.UseQueries())

	// di - inject dependencies
	c := controllers.NewControllers(
		s.db.UseQueries(),
		auth.NewUserSessionStore(),
		logging.NewDefaultLogger(),
	)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	// api
	mux.HandleFunc("/test-endpoint", h.HelloWorldHandler)
	mux.HandleFunc("/health", s.healthHandler)

	// auth forms
	mux.HandleFunc(routes.Login, c.Auth.MapLogin)
	mux.HandleFunc(routes.Register, c.Auth.MapRegister)
	mux.HandleFunc(routes.Logout, c.Auth.MapLogout)

	// pages
	// home
	mux.HandleFunc(routes.Home, c.Home.Map)
	// trips
	mux.HandleFunc(routes.Trips, c.Trips.Map)
	mux.HandleFunc(routes.NewTrips, c.Trips.MapNewTrips)
	// tripbuilder
	mux.HandleFunc(routes.TripBuilder, c.TripBuilder.Map)

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
