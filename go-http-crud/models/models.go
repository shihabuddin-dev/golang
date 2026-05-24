package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"username"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var Users = []User{}
