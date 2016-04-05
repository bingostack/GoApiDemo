package models

import (
	// "errors"
	"fmt"
	"log"
)

//Relationship
type Relationship struct {
	Id       int64  `json:"id"`
	Owner_id int64  `json:"-"`
	User_id  int64  `json:"user_id"`
	State    string `json:"state"`
	Type     string `json:"type"`
}

func init() {
	// syncdb
	if err := x.Sync(new(Relationship)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

// Get Relationship using owner_id and user_id
func getRelationship(owner_id int64, user_id int64) (Relationship, bool, error) {
	var relation Relationship
	relation.Owner_id = owner_id
	relation.User_id = user_id
	has, err := x.Get(&relation)
	return relation, has, err
}

// Update Relationship using fields
func updateRelationship(owner_id int64, user_id int64, state string, rtype string) error {
	_, err := x.Where("owner_id = ?", owner_id).Where("user_id = ?", user_id).
		Update(&Relationship{
			Owner_id: owner_id, User_id: user_id,
			State: state, Type: rtype})
	return err
}

// Update or Create Relationship handler
func NewRelationship(owner_id int64, user_id int64, state string) (
	Relationship, error) {
	var relation Relationship

	if state == "liked" {
		// if both like, then match
		vrelation, has, err := getRelationship(user_id, owner_id)
		if err != nil {
			log.Printf("Fail to get Relationship from database: %v\n", err)
			return relation, err
		}
		if has && vrelation.State == "liked" {
			state = "matched"
			vrelation.State = "matched"
			fmt.Printf("%v and %v matched!\n", owner_id, user_id)
			// update to "matched"
			_, err := x.Id(vrelation.Id).Cols("state").Update(&vrelation)
			if err != nil {
				log.Printf("Fail to create Relationship: %v\n", err)
				return relation, err
			}
		}
	}
	// Check if exists
	_, has, err := getRelationship(owner_id, user_id)
	if err != nil {
		log.Printf("Fail to get Relationship from database: %v\n", err)
		return relation, err
	}
	// Update or Insert
	if has {
		err = updateRelationship(owner_id, user_id, state, "relationship")
	} else {
		_, err = x.Where("owner_id = ?", owner_id).
			Where("user_id = ?", user_id).Insert(&Relationship{
				Owner_id: owner_id, User_id: user_id,
				State: state, Type: "relationship"})
	}
	if err != nil {
		log.Printf("Fail to create Relationship: %v\n", err)
		return relation, err
	}

	// get Relationship just added
	relation, _, err = getRelationship(owner_id, user_id)
	return relation, err
}

// Get Relationship handler
func GetRelationship(owner_id int64) ([]Relationship, error) {
	var relations []Relationship = make([]Relationship, 0)
	err := x.Where("owner_id = ?", owner_id).Find(&relations)
	if err != nil {
		log.Printf("Fail to get Relationship from database: %v\n", err)
		return nil, err
	}
	return relations, nil
}
