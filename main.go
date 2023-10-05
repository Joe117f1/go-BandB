package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT_NUMBER = ":8080"

func main() {
	router := mux.NewRouter()

	// Define routes for CRUD operations on users.
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/users/{%s}", id), GetUserById).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/users/{%s}", id), UpdateUser).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("/users/{%s}", id), DeleteUser).Methods("DELETE")

	// Start the HTTP server.
	http.Handle("/", router)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(PORT_NUMBER, nil)
}
