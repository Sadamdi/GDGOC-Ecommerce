package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// InitValidator menginisialisasi singleton instance validator
func InitValidator() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct DTO dan mengembalikan error format string jika ada
func ValidateStruct(s interface{}) error {
	if validate == nil {
		InitValidator()
	}

	err := validate.Struct(s)
	if err != nil {
		// Mengambil error pertama yang ditemui
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("field %s is invalid: %s", err.Field(), err.Tag())
		}
	}

	return nil
}
