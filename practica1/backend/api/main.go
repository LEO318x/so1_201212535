package main

import (
	"log"
	"net/http"

	"api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler2)
	r.HandleFunc("/test", routes.GetTestHandler).Methods("POST")
	r.HandleFunc("/sumar", routes.Sumar).Methods("POST")
	r.HandleFunc("/restar", routes.Restar).Methods("POST")
	r.HandleFunc("/multiplicar", routes.Multiplicar).Methods("POST")
	r.HandleFunc("/dividir", routes.Dividir).Methods("POST")
	log.Printf("Servidor en puerto 8000")
	http.ListenAndServe(":8000", r)

}
