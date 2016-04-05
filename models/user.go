package models

import (
	"errors"
	// "log"
)

//User
type User struct {
	Id   int64
	Name string `xorm:"unique"`
	Type string
}

func newUser(name string) error {
	_, err := x.Insert(&User{Name: name, Type: "user"})
	return err
}
