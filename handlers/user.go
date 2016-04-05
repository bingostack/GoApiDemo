package handlers

import (
	"GoApiDemo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetUser()
	fmt.Println(users)
	result, _ := json.Marshal(users)
	w.Write(result)
}
