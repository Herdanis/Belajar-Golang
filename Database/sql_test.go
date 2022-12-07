package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO customer(id,name,email,balance,rating,birth_date,married) VALUE('MNB','zxc','mnb@gmail.com',700000,70.3,'1993-10-01',false)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("name :", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("=============")
		fmt.Println("id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birth_date.Valid {
			fmt.Println("birth_date :", birth_date.Time)
		}
		fmt.Println("created_at :", created_at)
		fmt.Println("married :", married)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin"
	password := "admin"
	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Welcome username :", username)
	} else {
		fmt.Println("Username/Password Wrong")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin'; #"
	password := "admin"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Welcome back", username)
	} else {
		fmt.Println("Username/Password Wrong")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "poi"
	password := "mnb"
	script := "INSERT INTO user(username,password) VALUE(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	email := "tes@gmaill.com"
	comment := "test afdba rva ba"
	ctx := context.Background()
	script := "INSERT INTO comment(email,comment) VALUE(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success with id", insertId)
}

func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO comment(email,comment) VALUE(?,?)"
	statment, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statment.Close()
	for i := 0; i < 10; i++ {
		email := "qwerty" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment test ke " + strconv.Itoa(i)
		result, err := statment.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Data ke", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comment(email,comment) VALUE(?,?)"
	for i := 0; i < 10; i++ {
		email := "qwerty" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment test ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Data ke", id)
	}
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
