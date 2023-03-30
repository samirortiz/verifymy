package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AllUsers(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}
	users, err := allUsers()
	if len(users) == 0 {
		fmt.Fprintf(rw, "Error on GetUsers(): %v", "No users found\n")
		return
	}
	if err != nil {
		fmt.Fprintf(rw, "Error on GetUsers(): %v", err)
		return
	}
	json.NewEncoder(rw).Encode(users)
}

func UserById(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}
	user, err := userByID(r.URL.Query().Get("id"))
	if user == (User{}) {
		fmt.Fprintf(rw, "Error on UserById(): %v", "No user with specified id\n")
		return
	}
	if err != nil {
		fmt.Fprintf(rw, "Error on UserById(): %v", err)
		return
	}
	json.NewEncoder(rw).Encode(user)
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "Error on ParseForm(): %v", err)
		return
	}
	user, err := createUser(r.Form)
	if err != nil {
		fmt.Fprintf(rw, "Error on CreateUser(): %v", err)
		return
	}
	fmt.Fprintf(rw, "The last inserted user: %v\n", &user)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		fmt.Println("Method not allowed")
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "Error in UpdateUser(),ParseForm(): %v", err)
		return
	}
	var userId string = r.URL.Query().Get("id")
	user, err := updateUser(userId, r.Form)
	if err != nil {
		http.Error(rw, "No user to update", http.StatusNoContent)
		fmt.Fprintf(rw, "Error in UpdateUser(): %v", err)
		return
	}
	fmt.Fprintf(rw, "Updated user id: %v\n", &user)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		fmt.Println("Method not allowed")
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var userId, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(rw, "Error in DeleteUser(): %v", err)
		return
	}
	affected, err := deleteUser(userId)
	if err != nil {
		fmt.Fprintf(rw, "Error in DeleteUser(): %v", err)
		return
	}
	if affected != 0 {
		fmt.Fprintf(rw, "Deleted user id: %v\n", userId)
		return
	}
	fmt.Fprintf(rw, "Error in DeleteUser(): %v", "No user to be deleted\n")
}
