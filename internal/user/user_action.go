package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/fajar3108/lms-backend/pkg/helpers"
	"gorm.io/gorm"
)

type UserAction struct {
	db *gorm.DB
}

func NewUserAction(db *gorm.DB) *UserAction {
	return &UserAction{db: db}
}

func (a *UserAction) CreateUser(ctx context.Context, name, email, password string) (*User, error) {
	const operation = "user.action.create_user"

	hashed, err := helpers.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	userId, err := helpers.GenerateUUID()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	usr := &User{
		ID:       userId,
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	if err := a.db.WithContext(ctx).Create(usr).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, fmt.Errorf("%s: %w", operation, ErrEmailAlreadyExist)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return usr, nil
}

func (a *UserAction) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	const operation = "user.action.get_user_by_email"

	var usr User
	if err := a.db.WithContext(ctx).Where("email = ?", email).First(&usr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: %w", operation, ErrEmailNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	return &usr, nil
}

func (a *UserAction) GetUserByID(ctx context.Context, id string) (*User, error) {
	const operation = "user.action.get_user_by_id"

	var usr User
	if err := a.db.WithContext(ctx).Where("id = ?", id).First(&usr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: %w", operation, ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &usr, nil
}
