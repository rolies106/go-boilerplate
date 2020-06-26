package utils

import (
	"encoding/json"
	"net/http"
)

// Only return message and success
func Message(success bool, message string) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message}
}

// Return message with data
func MessageWithData(success bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message, "data": data}
}

// Return message, success, and also collection of errors (map)
func ErrorMessage(success bool, message string, errors interface{}) map[string]interface{} {
	return map[string]interface{}{"errors": errors, "message": message, "success": success}
}

// Directly respond collection into http
func Respond(w http.ResponseWriter, httpStatusCode int, data interface{}) {

	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Max-Age", "604800")
	// w.Header().Set("Access-Control-Allow-Headers", "x-requested-with, x-requested-by")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(&data)
}
