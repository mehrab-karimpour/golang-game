package service

import (
	"gameapp/repository"
	"gameapp/request/userrequest"
	"gameapp/response/userresponse"
)

type Service struct {
	repo repository.User
}

func (s Service) Register(req userrequest.RegisterRequest) (userresponse.RegisterResponse, error) {
	validationErrors := req.Validate()

	if validationErrors != nil {
		return userresponse.RegisterResponse{}, validationErrors
	}

	newUser, storeErr := s.repo.Store(req)
	if storeErr != nil {
		return userresponse.RegisterResponse{}, storeErr
	}

	return userresponse.RegisterResponse{User: newUser}, nil
}
