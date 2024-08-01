package model

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

// Ficha de informações necessárias e suas verificações
	type Token struct {
		ID     int    `json:"id"` // Será auto incrementando através do banco de dados
		Name   string `json:"name" validate:"required,min=3,max=40"`
		CPF    string `json:"cpf" validate:"required,len=11"`
		Email  string `json:"email" validate:"required,email"`
		Number string `json:"number" validate:"required,len=11"`
	}

var validate *validator.Validate

// Inicia o validator
func init() {
	validate = validator.New()
}

// Retorna para a struct o validator iniciado
func (token *Token) Validate() error {
	return validate.Struct(token)
}

// Função responsável pelos cadastros e envio para o banco de dados
func (token *Token) Create(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO userSignUp(name, cpf, email, number) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token.Name, token.CPF, token.Email, token.Number)
	return err
}
