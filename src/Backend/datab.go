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
func insertUser(password string, tempuser User) {
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
	_, err = db.Exec(sqlStatement, tempuser.Username, tempuser.Password, tempuser.Name)
	if err != nil {
		panic(err)
	}
	db.Close()
}

//getUser fetches user infomration from Database.
func getUser(password string, username string) {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT password,name FROM Users WHERE Username=$1`
	row := db.QueryRow(sqlStatement, username)
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
func DeleteUser(password string, username string) {
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
	_, err = db.Exec(sqlStatement, username)
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
	INSERT INTO events (id, date, type, username, description, title, priority)
	VALUES ($1, $2, (SELECT name FROM type WHERE name = $3), (SELECT username FROM users WHERE username= $4), $5, $6, $7)`
	_, err = db.Exec(sqlStatement, tempevent.ID, tempevent.Date, tempevent.Type, tempevent.UserName, tempevent.Description, tempevent.Title, tempevent.Priority)
	if err != nil {
		panic(err)
	}
	db.Close()
}

//RemoveEvent Function removes the event from the planner.
func RemoveEvent(password string, id int) {
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
	_, err = db.Exec(sqlStatement, id)
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
func FetchEventByID(password string, id int) event {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT * FROM events WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	var desc string
	var title string
	var priority string
	var tempid int
	var temptype string
	var date string
	var username string
	temp := event{}

	switch err := row.Scan(&tempid, &priority, &title, &desc, &username, &date, &temptype); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")

	case nil:
		fmt.Println(desc, title, priority, tempid, temptype, date, username)
		temp = event{
			ID:          tempid,
			Title:       title,
			Description: desc,
			Priority:    priority,
			Type:        temptype,
			Date:        date,
			UserName:    username,
		}
		return temp

	default:
		panic(err)
	}
	db.Close()
	return temp
}

//FetchEventsForUser all the events that have this property
func FetchEventsForUser(password string, source string, info string) event {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//just assume that the person doesn't search by ID.

	sqlStatement := `SELECT * FROM events WHERE $1 =$2`

	row := db.QueryRow(sqlStatement, source, info)

	var desc string
	var title string
	var priority string
	var tempid int
	var temptype string
	var tempdate string
	var username string
	temp := event{}

	switch err := row.Scan(&tempid, &priority, &title, &desc, &username, &tempdate, &temptype); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")

	case nil:
		fmt.Println(desc, title, priority, tempid, temptype, tempdate, username)
		temp = event{
			ID:          tempid,
			Title:       title,
			Description: desc,
			Priority:    priority,
			Type:        temptype,
			Date:        tempdate,
			UserName:    username,
		}
		return temp

	default:
		panic(err)
	}
	db.Close()
	return temp
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
	_, err = db.Exec(sqlStatement, "1", "Sport")
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
