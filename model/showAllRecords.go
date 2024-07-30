package model

import (
	"database/sql"
)

// Responsável por mostrar cadastros realizados
func GetAllTokens(db *sql.DB) ([]Token, error) {
	rows, err := db.Query("SELECT id, name, cpf, email, number FROM userSignUp")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []Token
	for rows.Next() {
		var token Token
		if err := rows.Scan(&token.ID, &token.Name, &token.CPF, &token.Email, &token.Number); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
