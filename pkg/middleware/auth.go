package middleware

import (
	"strings"

	errorhandler "github.com/fajar3108/lms-backend/pkg/error-handler"
	"github.com/fajar3108/lms-backend/pkg/token"
	"github.com/gofiber/fiber/v3"
)

func AuthRequired(jwtManager *token.JWTManager) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errorhandler.GlobalErrorHandler(c, fiber.NewError(fiber.StatusUnauthorized, "missing authorization header"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return errorhandler.GlobalErrorHandler(c, fiber.NewError(fiber.StatusUnauthorized, "authorization header must be in 'Bearer <token>' format"))
		}

		tokenStr := parts[1]
		claims, err := jwtManager.VerifyToken(tokenStr)
		if err != nil {
			return errorhandler.GlobalErrorHandler(c, fiber.NewError(fiber.StatusUnauthorized, "invalid or expired access token"))
		}

		c.Locals("userID", claims.UserID)

		return c.Next()
	}
}
