package main

import (
	"crudapi/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("welcome to CRUD API")
	r := mux.NewRouter()
	routes.RegisterPersonRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9001", r))
}
