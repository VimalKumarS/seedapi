package middleware

import (
	"encoding/json"
	"net/http"
)

// ErrorRes : struct
type ErrorRes struct {
	Error   bool        `json:"error"`
	Message interface{} `json:"message"`
}

// ErrorResponse : http error response
func ErrorResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	error := ErrorRes{
		true,
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}
