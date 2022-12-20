package routes

import (
	"microtwo/handlers/home"
	"microtwo/servers"

	"github.com/gorilla/mux"
)

// HomeRoutes defines the routes for the home page
func HomeRoutes(s *servers.HttpServer, r *mux.Router) {
	r.HandleFunc("/", home.HomeHandler(s)).Methods("GET")
}
