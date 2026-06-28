package database

import (
	"fmt"

	"github.com/fajar3108/lms-backend/internal/course"
	"github.com/fajar3108/lms-backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&user.RefreshToken{},
		&course.Course{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database schemas: %w", err)
	}
	return nil
}
