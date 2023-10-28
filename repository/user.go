package repository

import (
	"gameapp/entity"
	"gameapp/http/request/userrequest"
)

type User interface {
	Store(u userrequest.RegisterRequest) (entity.User, error)
	FirstWhere(col string, val any) (*entity.User, error)
}
