package routes

import (
	"encoding/json"
	"net/http"
)

func HomeHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hola mundo desde la api!")
}
