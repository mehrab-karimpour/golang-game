package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Connection *sql.DB
}

func (db Mysql) NewConnection() *sql.DB {
	//TODO implement me
	panic("implement me")
}

var mysqlConnection *sql.DB

func init() {
	var connError error
	mysqlConnection, connError = sql.Open("mysql", "root:root@/gameapp")
	pingErr := mysqlConnection.Ping()

	if connError != nil || pingErr != nil {
		panic(fmt.Errorf("mysql connection error ! %w , ping error %v", connError, pingErr))
	} else {
		fmt.Println("database connection is ok...")
	}

}

func NewConnection() *sql.DB {
	return mysqlConnection
}
