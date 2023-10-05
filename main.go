package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
