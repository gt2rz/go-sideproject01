package home

import (
	"encoding/json"
	"microtwo/servers"
	"net/http"
)

func HomeHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Welcome to the home page")
	}
}
