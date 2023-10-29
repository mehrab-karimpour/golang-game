package repository

import (
	"gameapp/app/entity"
	"gameapp/app/http/request/userrequest"
)

type User interface {
	Store(u userrequest.RegisterRequest) (entity.User, error)
	FirstWhere(col string, val any) (*entity.User, error)
}
