package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
	"gameapp/request/userrequest"
)

func checkUsersExists(db *sql.DB) {
	query := `CREATE TABLE users (  id int(11) NOT NULL, 
              phone_number varchar(20) NOT NULL, 
              first_name varchar(255) NOT NULL, 
              last_name varchar(255) NOT NULL, 
              password varchar(255) NOT NULL,
              created_at timestamp NOT NULL DEFAULT current_timestamp(),
              updated_at timestamp NOT NULL DEFAULT current_timestamp
               )
              ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;`
	_, err := db.Exec(query)
	if err != nil {
		return
	}
}

func (db Mysql) Store(u userrequest.RegisterRequest) (entity.User, error) {
	checkUsersExists(db.Connection)

	exec, err := db.Connection.Query("select  * from users")
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println(exec.Columns())

	err = exec.Close()
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println()
	return entity.User{}, nil
}
