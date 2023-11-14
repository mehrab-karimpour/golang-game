package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Validator interface {
	Validate(customMessage *string) *ValidationErrors
}

type ValidationErrors struct {
	Error *string
}

func (val ValidationErrors) Validate(req Validator, customMessage *string) *ValidationErrors {

	if customMessage != nil {
		val.Error = customMessage
		return &val
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("%s is %s", GetTag(req, err.StructField()), err.Tag())
			val.Error = &msg
			return &val
		}
	}

	return nil
}

func GetTag(req Validator, Field string) string {
	userType := reflect.TypeOf(req)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		jsonTag := field.Tag.Get("json")

		if Field == field.Name {
			return jsonTag
		}
	}
	return Field
}
