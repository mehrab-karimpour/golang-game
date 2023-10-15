package userservice

import (
	"github.com/mehrab-karimpour/golang-game/entity"
	"github.com/mehrab-karimpour/golang-game/request/userrequest"
	"github.com/mehrab-karimpour/golang-game/response/userresponse"
)

type Repository interface {
	Store(u userrequest.RegisterRequest) (entity.User, error)
}
type Service struct {
	repo Repository
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
