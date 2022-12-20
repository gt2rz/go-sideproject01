package routes

import (
	"microtwo/handlers/auth"
	"microtwo/servers"

	"github.com/gorilla/mux"
)

func AuthRoutes(s *servers.HttpServer, r *mux.Router) {
	r.HandleFunc("/auth/signup", auth.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/auth/login", auth.LogInHandler(s)).Methods("POST")
	r.HandleFunc("/auth/forgotpassword", auth.ForgotPasswordHandler(s)).Methods("POST")
	r.HandleFunc("/auth/resetpassword/{resetToken}", auth.ResetPasswordHandler(s)).Methods("PATCH")
}
