package main

import (
	"log"
	"net/http"

	"api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/test", routes.GetTestHandler).Methods("POST")
	log.Printf("Servidor en puerto 8000")
	http.ListenAndServe(":8000", r)

}
