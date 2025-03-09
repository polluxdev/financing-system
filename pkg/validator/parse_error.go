package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (v *ValidatorImpl) ParseErrors(err interface{}) []string {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	errorMessages := make([]string, 0)

	// Define error message templates for tags
	tagMessages := map[string]func(err validator.FieldError) string{
		"gt": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
		},
		"required": func(err validator.FieldError) string {
			return fmt.Sprintf("%s is required", err.Field())
		},
		"datetime": func(err validator.FieldError) string {
			return fmt.Sprintf("%s is invalid format, should be %s", err.Field(), err.Param())
		},
		"isValidEmail": func(err validator.FieldError) string {
			return fmt.Sprintf("%s address is invalid", err.Field())
		},
		"isValidPhoneNumber": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must start with 62xx", err.Field())
		},
	}

	// Add error messages based on tags
	for _, e := range exception {
		if handler, exists := tagMessages[e.Tag()]; exists {
			// Handle predefined tags
			errorMessages = append(errorMessages, handler(e))
		} else {
			// Default case for unhandled tags
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
	}

	return errorMessages
}
