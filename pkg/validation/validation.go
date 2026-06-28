package validation

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(s any) error {
	err := v.validator.Struct(s)
	if err == nil {
		return nil
	}

	return fiberValidationError(err)
}
