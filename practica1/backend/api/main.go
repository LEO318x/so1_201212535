package main

import (
	"log"
	"net/http"

	"api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	enableCORS(r)
	r.HandleFunc("/", routes.HomeHandler2)
	r.HandleFunc("/test", routes.GetTestHandler).Methods("POST")
	r.HandleFunc("/sumar", routes.Sumar).Methods("POST")
	r.HandleFunc("/restar", routes.Restar).Methods("POST")
	r.HandleFunc("/multiplicar", routes.Multiplicar).Methods("POST")
	r.HandleFunc("/dividir", routes.Dividir).Methods("POST")
	r.HandleFunc("/historial", routes.Historial).Methods("GET")
	r.HandleFunc("/historialhoy", routes.Historial2).Methods("GET")
	log.Printf("Servidor en puerto 8000")
	http.ListenAndServe(":8000", r)

}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}
