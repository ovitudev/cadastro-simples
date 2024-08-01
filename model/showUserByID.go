package model

import "database/sql"

// Função responsável por selecionar os campos da tabela e puxar do banco de dados para apresentar diretamente pelo postman (por exemplo)
func GetUserByID(db *sql.DB, id int) (*Token, error) {
	var token Token

	err := db.QueryRow("SELECT id, name, cpf, email, number FROM userSignUp WHERE id = ?", id).Scan(
		&token.ID, &token.Name, &token.CPF, &token.Email, &token.Number)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
