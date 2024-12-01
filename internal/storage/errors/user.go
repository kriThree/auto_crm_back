package storage_errors

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exist")
)
