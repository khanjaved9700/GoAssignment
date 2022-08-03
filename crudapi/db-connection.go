package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DataMigration() {
	var url = "root:javed@123@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed!...")
	}

	db.AutoMigrate(&Persons{})
}
