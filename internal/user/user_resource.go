package user

import (
	"time"
)

type UserResource struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResource struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    time.Time    `json:"expires_at"`
	User         UserResource `json:"user"`
}

func TransformUser(user *User) UserResource {
	return UserResource{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func TransformLogin(user *User, accessToken string, refreshToken string, expiresAt time.Time) LoginResource {
	return LoginResource{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		User:         TransformUser(user),
	}
}
