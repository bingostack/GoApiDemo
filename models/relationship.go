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

func NewRelationship(owner_id int64, user_id int64, state string) (
	Relationship, error) {
	var relation Relationship

	if state == "liked" {
		// if both like, then match
		var vrelation Relationship
		vrelation.Owner_id = user_id
		vrelation.User_id = owner_id
		_, err := x.Get(&vrelation)
		if err != nil {
			log.Printf("Fail to get Relationship from database: %v\n", err)
			return relation, err
		}
		if vrelation.State == "liked" {
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
	// Update or Insert
	_, err := x.Where("owner_id = ?", owner_id).Where("user_id = ?", user_id).
		Insert(&Relationship{
			Owner_id: owner_id, User_id: user_id,
			State: state, Type: "relationship"})
	if err != nil {
		log.Printf("Fail to create Relationship: %v\n", err)
		return relation, err
	}
	// get Relationship just added
	relation.Owner_id = owner_id
	relation.User_id = user_id
	_, err = x.Get(&relation)
	return relation, err
}

func GetRelationship(owner_id int64) ([]Relationship, error) {
	var relations []Relationship = make([]Relationship, 0)
	err := x.Where("owner_id = ?", owner_id).Find(&relations)
	if err != nil {
		log.Printf("Fail to get Relationship from database: %v\n", err)
		return nil, err
	}
	return relations, nil
}
