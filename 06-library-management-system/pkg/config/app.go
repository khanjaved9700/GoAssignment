package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//  connecting with database using gorm.. gorm can make table by self if there is no table availabe in the database
func Connect() {
	d, err := gorm.Open("mysql", "root:javed@123@/bookms_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("connection failed...%v", err)
	}
	db = d
}

//  send database conections to the requester...
func GetDB() *gorm.DB {
	return db
}
