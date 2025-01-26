package request

import "github.com/go-playground/validator/v10"

func ValidationErrorResponse(err error) map[string]string {
	validationErrors := map[string]string{}

	// Parse the validation errors
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			// Customize messages based on field and validation tag
			switch fieldErr.Tag() {
			case "required":
				validationErrors[fieldErr.Field()] = fieldErr.Field() + " is required"
			case "email":
				validationErrors[fieldErr.Field()] = "Invalid email format"
			case "password":
				validationErrors[fieldErr.Field()] = "Password must be at least 8 characters long, contain uppercase, lowercase, a number, and a special character"
			case "min":
				validationErrors[fieldErr.Field()] = fieldErr.Field() + " must be at least " + fieldErr.Param() + " characters long"
			default:
				validationErrors[fieldErr.Field()] = "Invalid value of " + fieldErr.Field() + " field"
			}
		}
	}
	return validationErrors
}
