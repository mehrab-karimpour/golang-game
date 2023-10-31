package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/app/entity"
	"gameapp/app/http/request/userrequest"
	"gameapp/pkg/richerror"
)

func (db Mysql) Store(u userrequest.RegisterRequest) (entity.User, error) {
	var newUser = entity.User{}

	queryResult, queryExecErr := db.Connection.Exec("insert into users(phone_number,first_name,last_name,password) values (?,?,?,?)",
		u.PhoneNumber, u.FirstName, u.LastName, u.Password)

	if queryExecErr != nil {
		fmt.Println("error ", queryExecErr)
		return newUser, richerror.New().
			Kind(richerror.KindUnexpected).
			Message(queryExecErr.Error())
	}

	currentRowId, _ := queryResult.LastInsertId()

	stmt := db.Connection.QueryRow("SELECT  * FROM users where id=? LIMIT 1", currentRowId)

	scanErr := stmt.Scan(&newUser.ID, &newUser.PhoneNumber, &newUser.FirstName,
		&newUser.LastName, &newUser.Password, &newUser.CreatedAt, &newUser.UpdatedAt)
	if scanErr != nil {
		return newUser, richerror.New().
			Operation("mysql.store").
			Kind(richerror.KindUnexpected).
			Message(queryExecErr.Error())
	}
	return newUser, nil
}
func (db Mysql) FirstWhere(col string, val any) (*entity.User, error) {

	stmt := db.Connection.QueryRow("select * from users where "+col+" = ?", val)
	user := entity.User{}
	scanErr := stmt.Scan(&user.ID, &user.PhoneNumber, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt,
		&user.UpdatedAt)

	if scanErr != nil && scanErr == sql.ErrNoRows {
		return nil, richerror.New().
			Operation("mysql.FirstWhere").
			Kind(richerror.KindUnexpected).
			Message(scanErr.Error())
	}

	return &user, nil

}
