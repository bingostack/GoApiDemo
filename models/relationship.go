package models

import (
	// "errors"
	"log"
)

//Relationship
type Relationship struct {
	Id       int64
	Owner_id int64
	User_id  int64
	State    string
	Type     string
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
