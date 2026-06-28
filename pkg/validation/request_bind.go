package validation

import (
	"errors"

	errorhandler "github.com/fajar3108/lms-backend/pkg/error-handler"
	"github.com/gofiber/fiber/v3"
)

func RequestBind[T any](c fiber.Ctx, req *T) error {
	err := c.Bind().Body(req)

	if err == nil {
		return nil
	}

	_, ok := err.(*errorhandler.ValidationError)

	if !ok {
		var e *fiber.Error
		if errors.As(err, &e) {
			return e
		}
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body payload")
	}

	return err
}
