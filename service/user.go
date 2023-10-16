package service

import (
	"database/sql"
	"fmt"
	"gameapp/repository"
	"gameapp/request/userrequest"
	"gameapp/response/userresponse"
)

type User struct {
	Repo repository.User
}

func (u User) Register(req userrequest.RegisterRequest) (userresponse.RegisterResponse, error) {
	var response userresponse.RegisterResponse

	validationErrors := req.Validate()
	if validationErrors != nil || !u.phoneNumberIsUnique(req.PhoneNumber) {

		return response, validationErrors
	}

	newUser, storeErr := u.Repo.Store(req)
	if storeErr != nil {
		fmt.Println("storeErr", storeErr)
		return response, storeErr
	}
	response.User = newUser

	return response, nil
}

func (u User) phoneNumberIsUnique(phoneNumber string) bool {
	row, err := u.Repo.FirstWhere("phone_number", phoneNumber)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("userAlreadyExistsErr", err)
		return true
	}
	return row.ID == 0
}
