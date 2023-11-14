package userrequest

import "gameapp/app/http/request"

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=11,max=11"`
	Password    string `json:"password" validate:"required,min=6,max=12"`
}

func (login LoginRequest) Validate(customMessage *string) *request.ValidationErrors {
	var validation request.ValidationErrors
	return validation.Validate(login, customMessage)
}
