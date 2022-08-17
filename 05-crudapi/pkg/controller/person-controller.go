package controller

import (
	"crudapi/pkg/models"
	"crudapi/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "gorm.io/gorm/utils"
)

// createing new global variable

var NewPerson models.Person

//  get all data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var newpersons = models.GetAllPerson()
	res, _ := json.Marshal(newpersons)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  fetch data by providing id
func GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pId := vars["pid"]
	// fmt.Printf("%v", bookId)
	ID, err := strconv.ParseInt(pId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
		// log.Fatal(err)
	}

	personDetails, _ := models.GetPersonById(ID)
	res, _ := json.Marshal(personDetails)

	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

//  create personn entery
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	createNewPerson := &models.Person{}
	utils.ParseBody(r, createNewPerson)
	p := createNewPerson.CreatePerson()
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  delete data by providing id
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pId := vars["pid"]

	ID, err := strconv.ParseInt(pId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing")
	}
	book := models.DeletePerson(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  update person by providing id and replace it
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var updatedPerson = &models.Person{}
	utils.ParseBody(r, updatedPerson)
	vars := mux.Vars(r)
	pId := vars["pid"]
	ID, err := strconv.ParseInt(pId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing")
	}

	personDetails, db := models.GetPersonById(ID)

	if updatedPerson.FirstName != "" {
		personDetails.FirstName = updatedPerson.FirstName
	}
	if updatedPerson.LastName != "" {
		personDetails.LastName = updatedPerson.LastName
	}
	if updatedPerson.Email != "" {
		personDetails.Email = updatedPerson.Email
	}
	db.Save(&personDetails)
	res, _ := json.Marshal(personDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
