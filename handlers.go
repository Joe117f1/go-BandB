package main

import (
	"errors"
	"net/http"
)

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	writeJsonResponse(w, newUser, http.StatusCreated)
}

// GetUsers returns a list of all users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get users logic
}

// GetUserById returns a specific user by ID.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := extractQueryParam(r, id)
	user, err := findUserById(userId)
	if err != nil {
		errorHandler(w, err, http.StatusNotFound)
		return
	}
	writeJsonResponse(w, user, http.StatusOK)
}

// UpdateUser handles updating an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Update user logic
}

// DeleteUser handles the deletion of a user by ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Delete user by ID logic
}
