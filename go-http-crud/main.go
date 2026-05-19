package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"username"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{Id: 1, Name: "hanki", Age: 30, Email: "hanki@example.com"},
	{Id: 2, Name: "panki", Age: 25, Email: "panki@example.com"},
	{Id: 3, Name: "danki", Age: 29, Email: "danki@example.com"},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)                  // root path get request
	mux.HandleFunc("GET /health", healthHandler)          // /health path get request
	mux.HandleFunc("GET /users", getUserHandler)          // /users path get request
	mux.HandleFunc("POST /createUser", createUserHandler) // /createUser path post request
	fmt.Println("Server is running on http://localhost:5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! This is a simple HTTP server.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK, the server is healthy!")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "Invalid User Info")
		return
	}
	newUser.Id = len(users) + 1
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
	fmt.Println(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
