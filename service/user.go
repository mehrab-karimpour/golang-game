package service

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
	"gameapp/http/request/userrequest"
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

func (u User) Register(req userrequest.RegisterRequest) (entity.User, error) {
	newUser, storeErr := u.Repo.Store(req)
	if storeErr != nil {
		fmt.Println("storeErr", storeErr)

		return newUser, storeErr
	}

	return newUser, nil
}

func (u User) phoneNumberIsUnique(phoneNumber string) bool {
	row, err := u.Repo.FirstWhere("phone_number", phoneNumber)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("userAlreadyExistsErr", err)
		return true
	}
	return row.ID == 0
}
