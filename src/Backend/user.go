package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//User struct
type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}
type users []User

//slice of users.
var userlist users

//get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userlist)
}

//Create user
func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter Username and Password")
	}
	json.Unmarshal(reqbody, &newUser)
	insertUser(password,newUser)
	userlist = append(userlist, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
