package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-http-crud/db"
	"go-http-crud/models"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! This is a simple HTTP server.")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK, the server is healthy!")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "Invalid User Info")
		return
	}

	Query := `INSERT INTO users (username, age, email) VALUES ($1, $2, $3) RETURNING id`
	err = db.DB.QueryRow(context.Background(), Query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "Error creating user")
		return
	}
	models.Users = append(models.Users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	Query := `SELECT id, username, age, email FROM users`
	rows, err := db.DB.Query(context.Background(), Query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "Error fetching users")
		return
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
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

func GetSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam) // id convert from string to int for query
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}
	Query := `SELECT id, username, age, email FROM users WHERE id =$1`
	var user models.User
	err = db.DB.QueryRow(context.Background(), Query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email)
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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}
	Query := `UPDATE users SET username=$1, age=$2, email=$3 WHERE id=$4 RETURNING id, username, age, email`
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user data")
		return
	}
	err = db.DB.QueryRow(context.Background(), Query, updatedUser.Name, updatedUser.Age, updatedUser.Email, id).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.Email)
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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}

	Query := `DELETE FROM users WHERE id=$1 RETURNING id`
	err = db.DB.QueryRow(context.Background(), Query, id).Scan(&id)
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
