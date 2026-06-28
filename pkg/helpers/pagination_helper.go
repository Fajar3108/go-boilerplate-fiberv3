package helpers

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type PaginationMeta struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	MaxPage int   `json:"max_page"`
}

func NewPaginationMeta[T any](ctx context.Context, db *gorm.DB, page, limit int) (*PaginationMeta, error) {
	var total int64

	if err := db.WithContext(ctx).Model(new(T)).Count(&total).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	maxPage := 0
	if total > 0 && limit > 0 {
		maxPage = int(math.Ceil(float64(total) / float64(limit)))
	}

	return &PaginationMeta{
		Page:    page,
		Limit:   limit,
		Total:   total,
		MaxPage: maxPage,
	}, nil
}

func NormalizePagination(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}

func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}
