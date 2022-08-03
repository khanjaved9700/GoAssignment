package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	var person []Persons
	w.Header().Set("Content-Type", "application/json")
	db.Find(&person)
	json.NewEncoder(w).Encode(person)

}

func getById(w http.ResponseWriter, r *http.Request) {
	var person Persons
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db.First(&person, params["id"])
	json.NewEncoder(w).Encode(person)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Persons
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&person)
	db.Create(&person)
	json.NewEncoder(w).Encode(person)
}

func updateById(w http.ResponseWriter, r *http.Request) {
	var person Persons
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db.First(&person, params["id"])
	_ = json.NewDecoder(r.Body).Decode(&person)
	db.Save(&person)
	json.NewEncoder(w).Encode(person)

}

func deleteById(w http.ResponseWriter, r *http.Request) {
	var person Persons
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db.Delete(&person, params["id"])
	json.NewEncoder(w).Encode("persone is deleted")
}
