package repository

import (
	"gameapp/entity"
	"gameapp/request/userrequest"
)

type User interface {
	Store(u userrequest.RegisterRequest) (entity.User, error)
}
