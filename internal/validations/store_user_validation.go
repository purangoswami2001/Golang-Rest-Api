// internal/validations/user_validation.go

package validations

import (
	"api/internal/models"
	"regexp"
)

// ValidateEmail checks if the provided email is valid
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidateUser checks if the user fields are valid
func ValidateUser(user models.User, isUpdate bool) ValidationErrors {
	var validationErrors ValidationErrors

	if user.FirstName == "" {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "first_name",
			Message: "First name is required",
		})
	}
	if user.LastName == "" {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "last_name",
			Message: "Last name is required",
		})
	}
	if user.Email == "" {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "email",
			Message: "Email is required",
		})
	} else if !ValidateEmail(user.Email) {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "email",
			Message: "Invalid email format",
		})
	}
	if !isUpdate && user.Password == "" {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "password",
			Message: "Password is required",
		})
	} else if len(user.Password) < 6 {
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "password",
			Message: "Password must be at least 6 characters",
		})
	}

	return validationErrors
}
