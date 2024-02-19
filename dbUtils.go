package main

import "errors"

// insertUser insert one user record to the db
func insertUser(user User) error {
	// insert to db logic...
	users = append(users, user)
	return nil
	// insert to db, might fail in reality
}

// findUserById retrieve a user by id from db, return error if not found
func findUserById(id string) (*User, error) {
	// db query logic...
	var searchedUser *User
	var err error

	for _, user := range users {
		if user.ID == id {
			searchedUser = &user
			err = nil
			break
		}
	}

	if searchedUser != nil {
		return searchedUser, nil
	}

	err = errors.New("couldn'd find user")
	return nil, err
}

// fetchUsers fetch all users from the db
func fetchUsers() (*Users, error) {
	// db query logic...
	return &users, nil
}
