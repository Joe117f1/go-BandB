package main

var users []User

func insertUser(user User) error {
	users = append(users, user)
	return nil
	// insert to db, might fail in reality
}
