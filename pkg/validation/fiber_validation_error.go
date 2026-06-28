package validation

import (
	"strings"

	errorhandler "github.com/fajar3108/lms-backend/pkg/error-handler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func fiberValidationError(err error) error {
	messages := make(map[string]string)

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	for _, e := range ve {
		messages[strings.ToLower(e.Field())] = e.Tag()
	}

	if len(messages) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Validation error")
	}

	return errorhandler.NewValidationError("Validation error", messages)
}
