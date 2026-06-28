package test_helper

import (
	"testing"

	"github.com/fajar3108/lms-backend/database"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(
		sqlite.Open("file::memory:?cache=shared&_foreign_keys=on"),
		&gorm.Config{
			TranslateError: true,
		},
	)
	require.NoError(t, err)
	require.NoError(t, db.Exec("PRAGMA foreign_keys = ON").Error)
	require.NoError(t, database.AutoMigrate(db))

	return db
}
