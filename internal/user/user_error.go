package user

import "errors"

var (
	ErrEmailNotFound     = errors.New("email not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exist")
)
