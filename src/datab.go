package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<USER DATABASE FUNCTIONS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//insertUser adds a user to the database with all the required information.
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
	db.Close()
}

//getUser fetches user infomration from Database.
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
	db.Close()
}

//DeleteUser deletes a user from database
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
	db.Close()
}

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<EVENT DATABASE FUNCTIONS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/

//AddEvent Function adds an event to the database.
func AddEvent(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO events (id, date, type_id, username, description, title, priority)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(sqlStatement, "1", "24.04.1998", "2", "{lazy}", "{I was born}", "{Birthday}", "urgent")
	if err != nil {
		panic(err)
	}
	db.Close()
}

//Remove EVENT Function

//EDIT EVENT Function

//Fetch Event Information Funciton.
