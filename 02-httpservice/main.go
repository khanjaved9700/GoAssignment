package main

import (
	"fmt"
	"httpservices/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()                                                // creating a new router
	r.HandleFunc("/wordcount", controller.HttpServices).Methods("Post") // register on request handler on it
	fmt.Println("servr is running at PORT 4004")
	log.Fatal(http.ListenAndServe(":4004", r))

}

/*The router is the main router for the application and willcpassed as
parameter to the server. It will receive all HTTP connections and pass it on to the
request handlers you will register on it.
*/
