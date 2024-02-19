package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func decodeRequestBody(req *http.Request, entity interface{}) error {
	err := json.NewDecoder(req.Body).Decode(entity)
	return err
}

func writeJsonResponse(w http.ResponseWriter, res interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func errorHandler(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func extractQueryParam(r *http.Request, param string) string {
	queryParams := mux.Vars(r)
	return queryParams[param]
}
