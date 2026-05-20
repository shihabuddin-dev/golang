package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

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

func connectDB() {
	var err error
	connectStr := "postgres://username:password@localhost:5432/database_name" // connection string
	db, err = pgx.Connect(context.Background(), connectStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database successfully!")
}

func main() {
	connectDB()
	defer db.Close(context.Background())

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /", rootHandler)                    // root path get request
	mux.HandleFunc("GET /health", healthHandler)            // /health path get request
	mux.HandleFunc("GET /users", getUserHandler)            // /users path get request
	mux.HandleFunc("POST /createUser", createUserHandler)   // /createUser path post request
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler) // /users/{id} path get request
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)    // /users/{id} path put request
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler) // /users/{id} path delete request

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

	Query := `INSERT INTO users (username, age, email) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(context.Background(), Query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "Error creating user")
		return
	}
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	Query := `SELECT id, username, age, email FROM users`
	rows, err := db.Query(context.Background(), Query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "Error fetching users")
		return
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(w, "Error scanning user data")
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam) // id convert from string to int for query
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}
	Query := `SELECT id, username, age, email FROM users WHERE id =$1`
	var user User
	err = db.QueryRow(context.Background(), Query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "User not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}
	Query := `UPDATE users SET username=$1, age=$2, email=$3 WHERE id=$4 RETURNING id, username, age, email`
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user data")
		return
	}
	err = db.QueryRow(context.Background(), Query, updatedUser.Name, updatedUser.Age, updatedUser.Email, id).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "User not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error updating user")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}

	Query := `DELETE FROM users WHERE id=$1 RETURNING id`
	err = db.QueryRow(context.Background(), Query, id).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "User not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting user")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", id)
}
