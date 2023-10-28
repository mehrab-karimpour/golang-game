package service

import (
	"database/sql"
	"fmt"
	"gameapp/http/request/userrequest"
	"gameapp/http/response/userresponse"
	"gameapp/repository"
)

type AuthUserInterface interface {
	AuthPrepare
	AuthGenerator
	AuthValidator
}

type AuthPrepare interface {
	New() Auth
}

type AuthGenerator interface {
	GenerateAccessToken(ID int) *string
	GenerateRefreshToken(ID int) *string
}

type AuthValidator interface {
	TokenIsValid(tokenStr string) bool
	GetAuth(tokenStr string) *any
}

type User struct {
	Repo     repository.User
	UserAuth AuthUserInterface
}

func (u User) Register(req userrequest.RegisterRequest) (userresponse.RegisterResponse, error) {
	var response userresponse.RegisterResponse
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
