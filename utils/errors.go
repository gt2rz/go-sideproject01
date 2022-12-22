package utils

import "errors"

// Signup errors
var ErrSignUpFailed = errors.New("sign up failed")
var ErrUserAlreadyExists = errors.New("sign up failed")

// Auth errors
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUnauthorizedAccess = errors.New("unauthorized access")
var ErrResetTokenExpired = errors.New("reset token expired")
var ErrResetTokenNotSaved = errors.New("reset token not saved")
var ErrGenerateRandomString = errors.New("error generating random string")

// Common errors
var ErrAnErrorOccurred = errors.New("an error has occurred")
var ErrInvalidRequest = errors.New("invalid request")

// User errors
var ErrUserNotFound = errors.New("user not found")
var ErrUserNotSaved = errors.New("user not saved")

// Customer errors
var ErrCustomerNotFound = errors.New("customer not found")
var ErrCustomerNotSaved = errors.New("customer not saved")
