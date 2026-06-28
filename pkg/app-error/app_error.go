package app_error

import (
	"errors"
	"fmt"
)

type AppError struct {
	Kind      Kind
	Code      string
	Message   string
	Operation string
	Err       error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Operation, e.Code, e.Err)
	}

	return fmt.Sprintf("%s: %s", e.Operation, e.Code)
}

func NewAppError(kind Kind, code, message, operation string, err error) error {
	return &AppError{
		Kind:      kind,
		Code:      code,
		Message:   message,
		Operation: operation,
		Err:       err,
	}
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func GetAppError(err error) (*AppError, bool) {
	var appErr *AppError

	if !errors.As(err, &appErr) {
		return nil, false
	}

	return appErr, true
}
