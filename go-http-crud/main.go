package main

import (
	"context"
	"fmt"
	"go-http-crud/db"
	"go-http-crud/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("env file not found")
	}

	db.ConnectDB()
	defer db.DB.Close(context.Background())

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /", handlers.RootHandler)                    // root path get request
	mux.HandleFunc("GET /health", handlers.HealthHandler)            // /health path get request
	mux.HandleFunc("GET /users", handlers.GetUserHandler)            // /users path get request
	mux.HandleFunc("POST /createUser", handlers.CreateUserHandler)   // /createUser path post request
	mux.HandleFunc("GET /users/{id}", handlers.GetSingleUserHandler) // /users/{id} path get request
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUserHandler)    // /users/{id} path put request
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUserHandler) // /users/{id} path delete request

	fmt.Println("Server is running on http://localhost:5000")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
