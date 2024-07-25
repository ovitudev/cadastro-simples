package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Inicializando o banco de dados
	DB, err = sql.Open("mysql", "root:281710@tcp(127.0.0.1:3306)/cadastrosimples")
	if err != nil {
		log.Fatal(err)
	}
}
