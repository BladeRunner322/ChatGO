package core_errors

import "errors"

var (
	ErrNotFound             = errors.New("not found")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrConflict             = errors.New("conflict")
	ErrInvalidToken         = errors.New("invalid token")
	ErrExpiredToken         = errors.New("token expired")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)
