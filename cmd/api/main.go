package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fajar3108/lms-backend/config"
	"github.com/fajar3108/lms-backend/database"
	"github.com/fajar3108/lms-backend/internal/router"
	errorhandler "github.com/fajar3108/lms-backend/pkg/error-handler"
	"github.com/fajar3108/lms-backend/pkg/validation"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fatal: failed to load configuration: %v", err)
	}

	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Fatal: failed to initialize database: %v", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Fatal: migration failed: %v", err)
	}

	validator := validation.NewValidator()

	app := fiber.New(fiber.Config{
		ErrorHandler:    errorhandler.GlobalErrorHandler,
		StructValidator: validator,
	})

	router.SetupRoutes(app, cfg, db, validator)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.AppPort
	}

	log.Println("Server is running on port " + port)

	if err := http.ListenAndServe(":"+port, adaptor.FiberApp(app)); err != nil {
		log.Fatal(err)
	}
}
