package services_errors

import "errors"

var (
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrNoAuthorizationTokens = errors.New("no authorization tokens")
	ErrIncorrectAuthToken = errors.New("incorrect auth token")
)
