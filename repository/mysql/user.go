package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
	"gameapp/request/userrequest"
)

func init() {
	checkUsersExists(mysqlConnection)
}

func checkUsersExists(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
id int(11) NOT NULL AUTO_INCREMENT,
phone_number varchar(15) NOT NULL,
first_name varchar(100) DEFAULT NULL,
last_name varchar(100) DEFAULT NULL,
password varchar(255) NOT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp(),
updated_at timestamp NOT NULL DEFAULT current_timestamp(),
 PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
 `
	_, execError := db.Exec(query)
	if execError != nil {
		fmt.Println("execError ", execError)
		return
	}
}

func (db Mysql) Store(u userrequest.RegisterRequest) (entity.User, error) {
	var newUser = entity.User{}

	queryResult, queryExecErr := db.Connection.Exec("insert into users(phone_number,first_name,last_name,password) values (?,?,?,?)",
		u.PhoneNumber, u.FirstName, u.LastName, u.Password)

	if queryExecErr != nil {
		fmt.Println("error ", queryExecErr)
		return newUser, queryExecErr
	}

	currentRowId, _ := queryResult.LastInsertId()

	stmt := db.Connection.QueryRow("SELECT  * FROM users where id=? LIMIT 1", currentRowId)

	scanErr := stmt.Scan(&newUser.ID, &newUser.PhoneNumber, &newUser.FirstName,
		&newUser.LastName, &newUser.Password, &newUser.CreatedAt, &newUser.UpdatedAt)
	if scanErr != nil {
		return newUser, scanErr
	}
	return newUser, nil
}
func (db Mysql) FirstWhere(col string, val any) (*entity.User, error) {
	stmt := db.Connection.QueryRow("select * from users where "+col+" = ?", val)
	user := entity.User{}
	scanErr := stmt.Scan(&user.ID, &user.PhoneNumber, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt,
		&user.UpdatedAt)
	return &user, scanErr

}
