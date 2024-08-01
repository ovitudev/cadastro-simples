package view

import (
	"encoding/json"
	"net/http"
)

// Padroniza o formato de resposta e requisição do código
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
