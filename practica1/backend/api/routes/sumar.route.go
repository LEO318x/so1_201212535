package routes

import (
	"api/controller"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Sumar(w http.ResponseWriter, r *http.Request) {
	var data models.Sumar
	json.NewDecoder(r.Body).Decode(&data)
	num1 := data.Num1
	num2 := data.Num2
	resultado := controller.RealizarSuma(num1, num2)
	log.Println(resultado)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"resultado": %d}`, resultado)
}
