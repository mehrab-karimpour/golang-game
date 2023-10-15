package mysql

import (
	"gameapp/entity"
	"gameapp/request/userrequest"
)

func (db Mysql) Store(u userrequest.RegisterRequest) (entity.User, error) {
	return entity.User{}, nil
}
