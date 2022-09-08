package main

import (
	"fmt"
	"jwt/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":50012"
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/refresh", handlers.Refresh)

	fmt.Printf("Server running at Port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
