package main

import (
	"encoding/json"
	"net/http"
)

func decodeRequestBody(req *http.Request, entity interface{}) error {
	err := json.NewDecoder(req.Body).Decode(entity)
	return err
}

func writeJsonResponse(w http.ResponseWriter, res interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(statusCode)
}

func errorHandler(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}
