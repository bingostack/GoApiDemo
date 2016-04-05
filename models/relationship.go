package models

import (
	// "errors"
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

func newRelationship(owner_id int64, user_id int64, state string) error {
	_, err := x.Insert(&Relationship{Owner_id: owner_id, User_id: user_id, State: state, Type: "user"})
	return err
}
