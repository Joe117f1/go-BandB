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
