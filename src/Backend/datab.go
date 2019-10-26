package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"time"

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
	INSERT INTO users (username, password, name)
	VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "SpiffyRaccoon", "test", "Aleks")
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
	sqlStatement := `SELECT password,name FROM Users WHERE Username=$1`
	row := db.QueryRow(sqlStatement, "lazy")
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
	_, err = db.Exec(sqlStatement, "lazy")
	if err != nil {
		panic(err)
	}
	db.Close()
}

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<EVENT DATABASE FUNCTIONS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/

//AddEvent Function adds an event to the database.
func AddEvent(password string, tempevent event) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO events (id, date, type_id, username, description, title, priority)
	VALUES ($1, $2, (SELECT id FROM type WHERE id = $3), (SELECT username FROM users WHERE username= $4), $5, $6, $7)`
	_, err = db.Exec(sqlStatement, tempevent.ID, time.Now(), tempevent.Type, tempevent.UserName, tempevent.Description, tempevent.Title, tempevent.Priority)
	if err != nil {
		panic(err)
	}
	db.Close()
}

//RemoveEvent Function removes the event from the planner.
func RemoveEvent(password string, tempevent event) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	DELETE FROM events
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, tempevent.ID)
	if err != nil {
		panic(err)
	}
	db.Close()
}

//EditEvent Function to update the events in the database.
func EditEvent(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
UPDATE events
SET title = $1
WHERE id = $2`
	_, err = db.Exec(sqlStatement, "Yay", "1")
	if err != nil {
		panic(err)
	}
	db.Close()
}

//FetchEvent Information Funciton.
func FetchEvent(password string) {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT description,title, priority FROM events WHERE id=$1`
	row := db.QueryRow(sqlStatement, "1")
	var desc string
	var title string
	var priority string

	switch err := row.Scan(&desc, &title, &priority); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(desc, title, priority)
	default:
		panic(err)
	}
	db.Close()
}

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<TYPE DATABASE FUNCTIONS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/

//AddType of event
func AddType(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO type (id, name)
	VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, "2", "lazy")
	if err != nil {
		panic(err)
	}
	db.Close()
}

//DeleteType function deletes the type of event from the database.
func DeleteType(password string) {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	DELETE FROM type
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, "2")
	if err != nil {
		panic(err)
	}
	db.Close()
}
