package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func insertUser(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO users (username, password, uname)
	VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "{lazy}", "{help}", "{Aleks}")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func getUser(password string) {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT password,uname FROM Users WHERE Username=$1`
	row := db.QueryRow(sqlStatement, "{lazy}")
	var pass string
	var name string

	switch err := row.Scan(&pass, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(pass, name)
	default:
		panic(err)
	}
	defer db.Close()
}

//DeleteUser  deletes a user from database
func DeleteUser(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	DELETE FROM users
	WHERE username = $1;`
	_, err = db.Exec(sqlStatement, "{lazy}")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
