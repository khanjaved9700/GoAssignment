package main

import "gorm.io/gorm"

type Persons struct {
	gorm.Model
	Name  string `json:"name"`
	Age   uint32 `json:"age"`
	Email string `json:"email"`
}
