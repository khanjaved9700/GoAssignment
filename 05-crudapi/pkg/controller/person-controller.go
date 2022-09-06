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

// var NewPerson models.Person

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Hello Api User...</~h1>
	<h2> This is Javed just go and check all the crud operatins...</h2></br>
	<h3>Thanks for checking...</h3>
	`))
}

//  get all data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var newpersons = models.GetAllPerson()
	res, err := json.Marshal(newpersons)
	if err != nil {
		print("Error while marshelling")
		// log.Fatal(err)
	}
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
		print("error while parsing")
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
	utils.ParseBody(r, createNewPerson) ///calling ParseBody functions for parsing Body
	p := createNewPerson.CreatePerson()
	res, err := json.Marshal(p)
	if err != nil {
		print("error while marshelling")
		// log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  delete data by providing id
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pId := vars["pid"]

	ID, err := strconv.ParseInt(pId, 0, 0)
	if err != nil {
		fmt.Printf("error while marshelling")
	}
	book := models.DeletePerson(ID)
	res, err := json.Marshal(book)
	if err != nil {
		print("error while marshelling")
		// log.Fatal(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  update person by providing id and replace it
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var updatedPerson = &models.Person{}
	utils.ParseBody(r, updatedPerson) //calling ParseBody functions for parsing Body
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
	res, err := json.Marshal(personDetails)
	if err != nil {
		print("error while marshelling")
		// log.Fatal(err)
	}
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
