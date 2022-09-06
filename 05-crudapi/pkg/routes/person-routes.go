package routes

import (
	"crudapi/pkg/controller"

	"github.com/gorilla/mux"
)

// register routes
var RegisterPersonRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/createperson/", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/getalldata/", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/getdata/{pid}", controller.GetPersonById).Methods("GET")
	router.HandleFunc("/updateperson/{pid}", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/deleteperson/{pid}", controller.DeletePerson).Methods("DELETE")

}
