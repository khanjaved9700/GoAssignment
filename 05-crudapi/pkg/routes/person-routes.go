package routes

import (
	"crudapi/pkg/controller"

	"github.com/gorilla/mux"
)

// register routes
var RegisterPersonRoutes = func(router *mux.Router) {
	router.HandleFunc("/creates/", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/creates/", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/creates/{pid}", controller.GetPersonById).Methods("GET")
	router.HandleFunc("/creates/{pid}", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/creates/{pid}", controller.DeletePerson).Methods("DELETE")

}
