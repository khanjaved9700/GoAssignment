package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/persons", getAll).Methods("GET")
	r.HandleFunc("/persons/{id}", getById).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/{id}", updateById).Methods("PUT")
	r.HandleFunc("/persons/{id}", deleteById).Methods("DELETE")

	fmt.Println("Server running at PORT 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
