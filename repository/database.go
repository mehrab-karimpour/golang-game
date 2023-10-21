package repository

import (
	"database/sql"
	"gameapp/repository/mysql"
)

const dbUsed = "mysql"

type DatabaseInterface interface {
	NewConnection() *sql.DB
}

type Database struct {
	Connection *sql.DB
}

var mysqlDatabase mysql.Mysql

func init() {
	if dbUsed == "mysql" {
		mysqlDatabase.Connection = mysql.NewConnection()
	}
}

func DataBaseService() *mysql.Mysql {
	return &mysqlDatabase
}
