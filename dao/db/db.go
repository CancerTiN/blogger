package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	driverName := "mysql"
	dataSourceName := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	DB = db
}
