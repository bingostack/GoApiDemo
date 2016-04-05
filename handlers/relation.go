package handlers

import (
	"GoApiDemo/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Relationship handler
func RelationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["owner_id"])
	owner_id, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	relations, err := models.GetRelationship(owner_id)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	fmt.Println(relations)
	result, _ := json.Marshal(relations)
	w.Write(result)
}

func AddRelationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["owner_id"])
	owner_id, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	fmt.Println(vars["user_id"])
	user_id, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}

	var relation models.Relationship
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&relation); err != nil {
		errorHandler(w, r, 500, err)
		return
	}

	relation, err = models.NewRelationship(owner_id, user_id, relation.State)
	if err != nil {
		errorHandler(w, r, 500, err)
		return
	}
	fmt.Println(relation)
	result, _ := json.Marshal(relation)
	w.Write(result)
}
