package main

import (
	"GoApiDemo/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.Methods("GET").Path("/users").HandlerFunc(handlers.UserHandler)
	r.Methods("POST").Path("/users").HandlerFunc(handlers.AddUserHandler)
	r.Methods("GET").Path("/users/{owner_id}/relationships").
		HandlerFunc(handlers.RelationHandler)
	r.Methods("PUT").Path("/users/{owner_id}/relationships/{user_id}").
		HandlerFunc(handlers.AddRelationHandler)
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
