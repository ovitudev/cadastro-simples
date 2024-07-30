package model

import (
	"database/sql"
)

func (token *Token) Delete(db *sql.DB) error {
	stmt, err := db.Prepare("TRUNCATE TABLE userSignUp")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return err
}
