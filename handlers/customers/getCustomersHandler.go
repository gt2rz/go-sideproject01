package customers

import (
	"context"
	"encoding/json"
	"microtwo/repositories"
	"microtwo/servers"
	"net/http"
)

type GetCustomerResponse struct {
	Customers []repositories.Customers `json:"customers"`
}

// GetCustomers returns all customers
func GetCustomersHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		customers, err := s.CustomerRepository.GetCustomersAll(context.Background())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
