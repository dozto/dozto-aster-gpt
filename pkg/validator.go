package pkg

import "github.com/go-playground/validator/v10"

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) error {
	errs := validate.Struct(data)

	if errs != nil {
		return errs
	}

	return nil
}

// func (v XValidator) TransErrorToResponse(err *error) []ErrorResponse {
// 	validationErrors := []ErrorResponse{}

// 	for _, err := range *err.(validator.ValidationErrors) {
// 		var elem ErrorResponse

// 		elem.FailedField = err.Field() // Export struct field name
// 		elem.Tag = err.Tag()           // Export struct tag
// 		elem.Value = err.Value()       // Export field value
// 		elem.Error = true

// 		validationErrors = append(validationErrors, elem)
// 	}

// 	return validationErrors
// }

var Valtor = &XValidator{
	validator: validate,
}
