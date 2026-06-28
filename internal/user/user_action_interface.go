package user

import "context"

//go:generate mockgen -source=$GOFILE -destination=../../test/mock/user/user_action_mock.go -package=user_mock
type UserActionInterface interface {
	CreateUser(ctx context.Context, name, email, password string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)
}
