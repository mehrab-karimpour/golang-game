package userrequest

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	FirstName   string `validate:"required" json:"first_name"`
	LastName    string `validate:"required" json:"last_name"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	Password    string `validate:"required" json:"password"`
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
