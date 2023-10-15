package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	connection *sql.DB
}

var mysqlConnection *sql.DB

func init() {
	var connError error
	mysqlConnection, connError = sql.Open("mysql", "root:root@/gameapp")
	if connError != nil {
		panic(fmt.Errorf("mysql connection error ! %w", connError))
	} else {
		fmt.Println("database connection is ok...")
	}
}

func NewConnection() *sql.DB {
	return mysqlConnection
}
