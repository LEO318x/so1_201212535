package routes

import (
	"api/controller"
	"encoding/json"
	"fmt"
	"net/http"
)

func Historial(w http.ResponseWriter, r *http.Request) {
	getHistorial, err := controller.GetHistorial()
	if err != nil {
		fmt.Println("ERR - Historial")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getHistorial)
}

func Historial2(w http.ResponseWriter, r *http.Request) {
	getHistorial, err := controller.GetHistorial2()
	if err != nil {
		fmt.Println("ERR - Historial")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getHistorial)
}
