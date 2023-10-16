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

func (s User) Register(req userrequest.RegisterRequest) (userresponse.RegisterResponse, error) {
	validationErrors := req.Validate()
	userAlreadyExists, userAlreadyExistsErr := s.Repo.FirstWhere("phone_number", req.PhoneNumber)

	if userAlreadyExistsErr != nil && userAlreadyExistsErr != sql.ErrNoRows {
		fmt.Println("userAlreadyExistsErr", userAlreadyExistsErr)
		return userresponse.RegisterResponse{}, userAlreadyExistsErr
	}

	if validationErrors != nil {
		return userresponse.RegisterResponse{}, validationErrors
	}

	if userAlreadyExists.ID == 0 {
		newUser, storeErr := s.Repo.Store(req)
		if storeErr != nil {
			fmt.Println("storeErr", storeErr)
			return userresponse.RegisterResponse{}, storeErr
		}
		return userresponse.RegisterResponse{User: newUser}, nil
	}
	return userresponse.RegisterResponse{}, nil
}
