package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//events slice assigning values.
var events = allEvents{}

//password for the database.
var password string

//database information
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Planner"
)

//Get Home Page
func homelink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	fmt.Scan(&password)
	DBConnecting()
	routing()
}

//DBConnecting connects the database and the server
func DBConnecting() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//routing handles all the routing that the user requests.
func routing() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homelink)
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getSingleEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	router.HandleFunc("/adduser", addUser).Methods("POST")
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
