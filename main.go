package main

import (
	"github.com/gorilla/mux"
	"net/http"
        "fmt"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
        fmt.Println("Now listening on http://localhost:8000...")
	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
