package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CustomPasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Minimum length of 8 characters
	if len(password) < 8 {
		return false
	}

	// At least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// At least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// At least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	// At least one special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}
