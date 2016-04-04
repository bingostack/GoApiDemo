package main

import (
	"github.com/gorilla/mux"
	// "log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write("Hello World!\n")
}