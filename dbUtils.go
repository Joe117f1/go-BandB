package main

import "errors"

func insertUser(user User) error {
	users = append(users, user)
	return nil
	// insert to db, might fail in reality
}

func findUserById(id string) (*User, error) {
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
