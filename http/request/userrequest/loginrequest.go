package userrequest

import "gameapp/http/request"

type LoginRequest struct {
	PhoneNumber string `validate:"required,min=11,max=11" json:"phone_number"`
	Password    string `validate:"required,min=6,max=12" json:"password"`
}

func (login LoginRequest) Validate(customMessage *string) *request.ValidationErrors {
	var validation request.ValidationErrors
	return validation.Validate(login, customMessage)
}
