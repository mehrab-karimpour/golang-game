package repository

import (
	"database/sql"
	"fmt"
	"gameapp/repository/mysql"
)

const dbUsed = "mysql"

type DatabaseInterface interface {
	NewConnection() *sql.DB
}

type Database struct {
	connection *sql.DB
}

var mysqlDatabase mysql.Mysql

func init() {
	if dbUsed == "mysql" {
		fmt.Println(mysql.NewConnection())
		mysqlDatabase.Connection = mysql.NewConnection()
	}
}

func DataBaseService() *mysql.Mysql {
	return &mysqlDatabase
}
