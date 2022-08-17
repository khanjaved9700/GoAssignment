package models

import (
	"crudapi/pkg/config"

	"github.com/jinzhu/gorm"
)

//  decleare global variabl

var db *gorm.DB

//  create struct for persons details
type Person struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// initialized function

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Person{})
}

//  create person

func (p *Person) CreatePerson() *Person {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

// get all person is type of slice becase all data store in the form slice..
func GetAllPerson() []Person {
	var persons []Person
	db.Find(&persons)
	return persons
}

//  get person by id
func GetPersonById(id int64) (*Person, *gorm.DB) {
	var getPerson Person
	db := db.Where("ID=?", id).Find(&getPerson)
	return &getPerson, db

}

// delete person
func DeletePerson(id int64) Person {
	var dltperson Person
	db.Where("ID=?", id).Delete(dltperson)
	return dltperson

}
