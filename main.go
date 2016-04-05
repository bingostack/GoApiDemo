package main

import (
	// "github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main() {
	r := NewRouter()
	http.Handle("/", r)
	fmt.Println("Now listening on http://localhost:8000...")
	http.ListenAndServe(":8000", r)
}

// func NewRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", HomeHandler)
// 	// r.HandleFunc("/users", m.)
// 	return r
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World!\n"))
// }
