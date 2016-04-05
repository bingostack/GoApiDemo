package handlers

import (
	"GoApiDemo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

//User handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUser()
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	fmt.Println(users)
	result, _ := json.Marshal(users)
	w.Write(result)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		errorHandler(w, r, 500, err)
		return
	}

	user, err := models.NewUser(user.Name)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	fmt.Println(user)
	result, _ := json.Marshal(user)
	w.Write(result)
}
