package handlers

import (
	"net/http"
)

// error handler
func errorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(`{"message": "Error"}`))
}
