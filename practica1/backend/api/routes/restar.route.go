package routes

import (
	"api/controller"
	"api/models"
	"encoding/json"
	"net/http"
)

func Restar(w http.ResponseWriter, r *http.Request) {
	var data models.Operacion
	json.NewDecoder(r.Body).Decode(&data)
	num1 := data.Num1
	num2 := data.Num2
	resultado := controller.RealizarResta(num1, num2)
	//log.Println(resultado)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
