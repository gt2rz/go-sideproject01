package routes

import (
	"microtwo/servers"

	"github.com/gorilla/mux"
)

func SetRoutes(s *servers.HttpServer, r *mux.Router) {
	// Add auth routes
	AuthRoutes(s, r)

	// Add home routes
	HomeRoutes(s, r)

	// Add customer routes
	CustomerRoutes(s, r)
}
