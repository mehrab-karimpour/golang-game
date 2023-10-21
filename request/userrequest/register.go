package userrequest

import (
	"gameapp/request"
)

type RegisterRequest struct {
	FirstName   string `validate:"required,min=3" json:"first_name"`
	LastName    string `validate:"required,min=3" json:"last_name"`
	PhoneNumber string `validate:"required,min=3" json:"phone_number"`
	Password    string `validate:"required,min=3" json:"password"`
}

func (req RegisterRequest) Validate(customMessage *string) *request.ValidationErrors {
	var validationErrors request.ValidationErrors
	return validationErrors.Validate(req, customMessage)
}
