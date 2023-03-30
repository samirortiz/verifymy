package main

type User struct {
	Id       int64  `json:"ID"`
	Name     string `json:"Name"`
	Age      string `json:"Age"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Address  string `json:"Address"`
}
