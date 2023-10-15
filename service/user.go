package service

import (
	"gameapp/repository"
	"gameapp/request/userrequest"
	"gameapp/response/userresponse"
)

type User struct {
	Repo repository.User
}

func (s User) Register(req userrequest.RegisterRequest) (userresponse.RegisterResponse, error) {
	validationErrors := req.Validate()

	if validationErrors != nil {
		return userresponse.RegisterResponse{}, validationErrors
	}

	newUser, storeErr := s.Repo.Store(req)
	if storeErr != nil {
		return userresponse.RegisterResponse{}, storeErr
	}

	return userresponse.RegisterResponse{User: newUser}, nil
}
