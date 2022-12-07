package golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectDb(t *testing.T) {
	db, err := sql.Open("mysql", "docker:123@tcp(127.0.0.1:3306)/docker")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
