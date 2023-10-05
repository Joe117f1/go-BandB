package main

import (
	"errors"
	"fmt"
	"net/http"
)

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser run")
	var newUser User

	err := decodeRequestBody(r, &newUser)

	if err != nil {
		errorHandler(w, err, http.StatusBadRequest)
		return
	}

	err = insertUser(newUser)
	if err != nil {
		errorHandler(w, errors.New("could not save data"), http.StatusInternalServerError)
		return
	}

	fmt.Println("users: ", users)
	writeJsonResponse(w, newUser, http.StatusCreated)
}

// GetUsers returns a list of all users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get users logic
}

// GetUser returns a specific user by ID.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get user by ID logic
}

// UpdateUser handles updating an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Update user logic
}

// DeleteUser handles the deletion of a user by ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Delete user by ID logic
}
