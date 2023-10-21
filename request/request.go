package request

import (
	"fmt"
	"gameapp/response"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(customMessage *string) *ValidationErrors
}

type ValidationErrors struct {
	Errors response.Messages
}

func (val ValidationErrors) Validate(req Validator, customMessage *string) *ValidationErrors {
	var messages = response.Messages{}
	if customMessage != nil {
		messages[fmt.Sprintf(*customMessage)] = fmt.Sprintf(*customMessage)
		val.Errors = messages
		return &val
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			messages[err.StructField()] = fmt.Sprintf("%s is %s", err.StructField(), err.Tag())
		}
		val.Errors = messages
		return &val
	}
	return nil

}
