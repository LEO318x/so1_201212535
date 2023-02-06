package routes

import (
	"api/models"
	"encoding/json"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Test!")
}

func GetTestHandler(w http.ResponseWriter, r *http.Request) {
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)
	log.Println(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK!")

}
