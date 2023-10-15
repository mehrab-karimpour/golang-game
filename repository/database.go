package repository

import (
	"database/sql"
	"gameapp/repository/mysql"
)

const dbUsed = "mysql"

type Database struct {
	connection *sql.DB
}

var databaseType *sql.DB

func init() {
	if dbUsed == "mysql" {
		databaseType = mysql.NewConnection()
	}
}

func DataBaseInstance() *sql.DB {
	return databaseType
}

type DatabaseInterface interface {
	NewConnection() *sql.DB
}
