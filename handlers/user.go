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
