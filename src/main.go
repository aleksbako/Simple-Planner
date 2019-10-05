package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//events slice assigning values.
var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to golang",
		Description: "test",
		Type:        "test type",
		Priority:    "urgent",
	},
}

//Get Home Page
func homelink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
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
