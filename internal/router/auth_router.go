package router

import (
	"github.com/fajar3108/lms-backend/internal/auth"
	"github.com/fajar3108/lms-backend/internal/user"
	"github.com/fajar3108/lms-backend/pkg/middleware"
	"github.com/fajar3108/lms-backend/pkg/token"
	"github.com/fajar3108/lms-backend/pkg/validation"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthRouter(route fiber.Router, jwtManager *token.JWTManager, db *gorm.DB, validator *validation.Validator) {
	authActions := auth.NewAuthAction(db)
	userAction := user.NewUserAction(db)
	authService := auth.NewAuthService(authActions, userAction, jwtManager)
	authCtrl := auth.NewAuthController(authService, validator)

	auth := route.Group("/auth")
	auth.Post("/register", authCtrl.Register)
	auth.Post("/login", authCtrl.Login)
	auth.Post("/refresh", authCtrl.Refresh)

	auth.Get("/me", middleware.AuthRequired(jwtManager), authCtrl.Me)
}
