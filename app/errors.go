package app

import (
	u "mortred/utils"
	"net/http"
)

// Return basic not found json format
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, http.StatusNotFound, u.Message(false, "This resources was not found on our server"))
	return
}