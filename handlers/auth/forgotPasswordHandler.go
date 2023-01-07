package auth

import (
	"context"
	"encoding/json"
	"errors"
	"microtwo/servers"
	"microtwo/services/mail/mailables"
	"microtwo/utils"
	"net/http"
	"net/smtp"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

type loginAuth struct {
	username, password string
}

// LoginAuth returns an smtp.Auth that implements the LOGIN authentication.	
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

// Start starts the SMTP login authentication process.
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

// Next is called after the server has sent a challenge.
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown fromServer")
		}
	}
	return nil, nil
}

// ForgotPasswordHandler handles the forgot password request
func ForgotPasswordHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = ForgotPasswordRequest{}

		// Decode the request body into the struct and failed if any error occur
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		// Get the user by email
		user, err := s.UserRepository.GetUserByEmail(context.Background(), request.Email)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Check if user is not found
		if user == nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidCredentials, http.StatusUnauthorized)
			return
		}

		// Generate a reset token
		resetToken, err := s.UserRepository.GenerateResetToken(context.Background(), user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create the mail data
		mailData := mailables.NewResetPasswordEmail(user.Email, resetToken)

		// Send the reset token to the user's email
		err = s.EmailService.SendResetPasswordEmail(mailData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the response
		utils.SendHttpResponse(w, &ForgotPasswordResponse{
			Message: "You will receive a reset email if user with that email exist",
		}, http.StatusOK)
	}
}
