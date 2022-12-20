package auth

import (
	"context"
	"encoding/json"
	"microtwo/servers"
	"microtwo/utils"
	"os"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Token     string `json:"token"`
	Status    bool   `json:"status"`
	Message   string `json:"message"`
}

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func LogInHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = LogInRequest{}

		// Decode the request body into the struct and failed if any error occur
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidCredentials, http.StatusBadRequest)
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

		// Compare the password with the hash
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidCredentials, http.StatusUnauthorized)
			return
		}

		// Create a new JWT token string
		JwtExpiresIn, _ := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
		claims := AppClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * JwtExpiresIn).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Create a response
		utils.SendHttpResponse(w, LogInResponse{
			Id:        user.Id,
			Email:     user.Email,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Phone:     user.Phone,
			Token:     signedToken,
			Status:    true,
			Message:   "Login successful",
		}, http.StatusOK)

	}
}
