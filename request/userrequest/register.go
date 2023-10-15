package userrequest

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	FirstName   string `validate:"required"`
	LastName    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Password    string `validate:"required"`
}

func (req RegisterRequest) Validate() error {
	validate := validator.New()
	phoneNumberIsUnique, err := req.checkPhoneNumberIsUnique()
	if !phoneNumberIsUnique && err != nil {
		return fmt.Errorf("unexpexted error : %w", err)
	}
	return validate.Struct(req)
}

func (req RegisterRequest) checkPhoneNumberIsUnique() (bool, error) {
	return true, nil
}
