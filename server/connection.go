package server

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "!QAZ2wsx#EDC",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "jwt",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, _ := sql.Open("mysql", config.FormatDSN())

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	Db = db
}
