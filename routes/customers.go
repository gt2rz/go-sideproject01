package routes

import (
	"microtwo/handlers/customers"
	"microtwo/servers"

	"github.com/gorilla/mux"
)

// CustomerRoutes defines the routes for the customers
func CustomerRoutes(s *servers.HttpServer, r *mux.Router) {
	r.HandleFunc("/customers", customers.GetCustomersHandler(s)).Methods("GET")
	r.HandleFunc("/customers/{id}", customers.GetCustomerByIdHandler(s)).Methods("GET")
}
