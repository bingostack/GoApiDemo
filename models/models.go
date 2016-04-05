package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var x *xorm.Engine

// create database engine
func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./goapidemo.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
}
