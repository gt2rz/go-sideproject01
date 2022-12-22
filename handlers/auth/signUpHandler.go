package auth

import (
	"context"
	"encoding/json"
	"microtwo/models"
	"microtwo/servers"
	"microtwo/utils"
	"time"

	"net/http"

	"github.com/google/uuid"
)

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}

type SignUpResponse struct {
	Status  bool   `json:"status"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

func SignUpHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}

		// Decode the request body into the struct and failed if any error occur
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrSignUpFailed, http.StatusBadRequest)
			return
		}

		// Check if the user already exists
		_, err = s.UserRepository.GetUserByEmail(context.Background(), request.Email)
		if err != utils.ErrUserNotFound {
			utils.SendHttpResponseError(w, utils.ErrUserAlreadyExists, http.StatusBadRequest)
			return
		}

		// Generate from the password with a predefined cost
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrSignUpFailed, http.StatusInternalServerError)
			return
		}

		// Generate a new UUID
		id, err := uuid.NewRandom()
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrSignUpFailed, http.StatusInternalServerError)
			return
		}

		// Create a new user
		var user = models.User{
			Id:        id.String(),
			Email:     request.Email,
			Password:  string(hashedPassword),
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Phone:     request.Phone,
			Verified:  false,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		// Save user to database
		err = s.UserRepository.SaveUser(context.Background(), user)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrSignUpFailed, http.StatusInternalServerError)
			return
		}

		// Return a 201 created status code
		utils.SendHttpResponse(w, SignUpResponse{
			Id:      user.Id,
			Message: "User created successfully",
			Status:  true,
		}, http.StatusCreated)

	}
}

// func sendEmailVerificationLink() {
// 	// Send email verification link to the user
// }

// func sendPhoneVerificationCode() {
// 	// Send phone verification code to the user
// }

// func sendWelcomeEmail() {
// 	// Send welcome email to the user
// }

// func sendWelcomeSMS() {
// 	// Send welcome SMS to the user
// }

// func sendWelcomePushNotification() {
// 	// Send welcome push notification to the user
// }
