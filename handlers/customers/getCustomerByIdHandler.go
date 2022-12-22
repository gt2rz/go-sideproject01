package customers

import (
	"context"
	"encoding/json"
	"microtwo/models"
	"microtwo/servers"
	"microtwo/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type GetCustomerByIdResponse struct {
	Customer models.Customer `json:"customer"`
}

// GetCustomers returns all customers
func GetCustomerByIdHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		customer, err := s.CustomerRepository.GetCustomerById(context.Background(), params["id"])

		if err == utils.ErrCustomerNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
