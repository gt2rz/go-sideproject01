package utils

import "errors"

// Signup errors
var ErrSignUpFailed = errors.New("sign up failed")
var ErrUserAlreadyExists = errors.New("sign up failed")

// Auth errors
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUnauthorizedAccess = errors.New("unauthorized access")

// Common errors
var ErrAnErrorOccurred = errors.New("an error has occurred")
var ErrInvalidRequest = errors.New("invalid request")
