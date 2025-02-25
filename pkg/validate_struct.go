package pkg

import (
	"fmt"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(d any) error {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		var dataErrorsValidation apierror.ValidationErrors
		errors, _ := err.(validator.ValidationErrors)
		for _, err := range errors {
			dataErrorsValidation = append(dataErrorsValidation, apierror.ValidationError{
				Field:      err.Field(),
				ErrorField: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()),
				Tag:        err.Tag(),
				Value:      err.Value(),
				Constraint: err.Param(),
			})
		}

		return dataErrorsValidation
	}

	return nil
}
