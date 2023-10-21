package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(customMessage *string) *ValidationErrors
}

type ValidationErrors struct {
	Errors []string
}

func (val ValidationErrors) Validate(req Validator, customMessage *string) *ValidationErrors {

	if customMessage != nil {
		val.Errors = append(val.Errors, fmt.Sprintf(*customMessage))
		return &val
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			val.Errors = append(val.Errors, fmt.Sprintf("%s is %s", err.StructField(), err.Tag()))
		}
		return &val
	}
	return nil

}
