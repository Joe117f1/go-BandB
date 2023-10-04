package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct represents a user entity.
type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func main() {
	router := mux.NewRouter()

	// Define routes for CRUD operations on users.
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Start the HTTP server.
	http.Handle("/", router)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser run")
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertUser(newUser)
	fmt.Println("users: ", users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
	w.WriteHeader(http.StatusCreated)
}
