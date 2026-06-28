package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fajar3108/lms-backend/internal/user"
	"gorm.io/gorm"
)

type AuthAction struct {
	db *gorm.DB
}

func NewAuthAction(db *gorm.DB) *AuthAction {
	return &AuthAction{db: db}
}

func (a *AuthAction) CreateRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) (*user.RefreshToken, error) {
	const operation = "auth.action.create_refresh_token"

	refreshToken := &user.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		IsRevoked: false,
	}

	if err := a.db.WithContext(ctx).Create(refreshToken).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return refreshToken, nil
}

func (a *AuthAction) RotateRefreshToken(
	ctx context.Context,
	userID string,
	oldToken string,
	newToken string,
	newExpiresAt time.Time,
) error {
	const operation = "auth.action.rotate_refresh_token"

	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.
			Model(&user.RefreshToken{}).
			Where(
				"token = ? AND user_id = ? AND is_revoked = ?",
				oldToken,
				userID,
				false,
			).
			Updates(map[string]any{
				"is_revoked": true,
			})

		if result.Error != nil {
			return fmt.Errorf(
				"%s: revoke old token: %w",
				operation,
				result.Error,
			)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf(
				"%s: %w",
				operation,
				ErrRefreshTokenRevoked,
			)
		}

		newRefreshToken := &user.RefreshToken{
			UserID:    userID,
			Token:     newToken,
			ExpiresAt: newExpiresAt,
			IsRevoked: false,
		}

		if err := tx.Create(newRefreshToken).Error; err != nil {
			return fmt.Errorf(
				"%s: create new token: %w",
				operation,
				err,
			)
		}

		return nil
	})
}

func (a *AuthAction) FindRefreshToken(ctx context.Context, token string) (*user.RefreshToken, error) {
	const operation = "auth.action.find_refresh_token"

	var rt user.RefreshToken
	if err := a.db.WithContext(ctx).Preload("User").Where("token = ?", token).First(&rt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: %w", operation, ErrRefreshTokenNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	return &rt, nil
}

func (a *AuthAction) RevokeAllUserRefreshTokens(ctx context.Context, userID string) error {
	const operation = "auth.action.revoke_all_user_refresh_tokens"

	tx := a.db.WithContext(ctx).
		Model(&user.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("is_revoked", true)

	if err := tx.Error; err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
