package routes

import (
	"api/controller"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Dividir(w http.ResponseWriter, r *http.Request) {
	var data models.Sumar
	json.NewDecoder(r.Body).Decode(&data)
	num1 := data.Num1
	num2 := data.Num2
	resultado := controller.RealizarDivision(num1, num2)
	log.Println(resultado)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"resultado": %s}`, resultado)
}
