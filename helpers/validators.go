package helpers

import "github.com/go-playground/validator/v10"

func HandleValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, fieldErr := range validationErrors {
		switch fieldErr.Tag() {
		case "required":
			errs[fieldErr.Field()] = fieldErr.Field() + " is required"
		case "min":
			errs[fieldErr.Field()] = fieldErr.Field() + " must be at least " + fieldErr.Param() + " characters long"
		case "alphanum":
			errs[fieldErr.Field()] = fieldErr.Field() + " must be alphanumeric"
		case "eqfield":
			errs[fieldErr.Field()] = fieldErr.Field() + " should be equal to " + fieldErr.Param()
		default:
			errs[fieldErr.Field()] = "Validation failed for " + fieldErr.Field()
		}
	}

	return errs
}
