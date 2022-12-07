package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "docker:123@tcp(127.0.0.1:3306)/docker?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
