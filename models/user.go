package models

import (
	// "errors"
	"log"
)

//User
type User struct {
	Id   int64
	Name string `xorm:"unique"`
	Type string
}

func init() {
	// 同步
	if err := x.Sync(new(User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func NewUser(name string) error {
	_, err := x.Insert(&User{Name: name, Type: "user"})
	return err
}

func GetUser() ([]User, error) {
	var user []User = make([]User, 0)
	err := x.Find(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
