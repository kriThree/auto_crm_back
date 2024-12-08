package storage_errors

import "errors"

var (
	ErrAdminRoleNotFound   = errors.New("admin role not found")
	ErrClientRoletNotFound = errors.New("client role not found")
	ErrOwnerRoleNotFound   = errors.New("owner role not found")
)
