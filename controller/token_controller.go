package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ovitudev/cadastro-simples.git/db"
	"github.com/ovitudev/cadastro-simples.git/model"
	"github.com/ovitudev/cadastro-simples.git/view"
)

// Função responsável pelos cadastros
func SignUp(w http.ResponseWriter, r *http.Request) {
	var newToken model.Token
	err := json.NewDecoder(r.Body).Decode(&newToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

// Validação básica do que é obrigátorio nos campos de cadastro
	if err := newToken.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	err = newToken.Create(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.JSONResponse(w, http.StatusOK, newToken)
}

func GetTokens(w http.ResponseWriter, r *http.Request) {
	tokens, err := model.GetAllSignUp(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.JSONResponse(w, http.StatusOK, tokens)
}
