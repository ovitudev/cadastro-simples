package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ovitudev/cadastro-simples.git/controller"
	"github.com/ovitudev/cadastro-simples.git/db"
)

// Função responsável pelas rotas e metodos utilizados nela
func main() {
	db.InitDB()
	router := mux.NewRouter()
	log.Println("Executando ...")

	router.HandleFunc("/cadastro", controller.SignUp).Methods("POST") // Rota para cadastros
	router.HandleFunc("/lista", controller.GetTokens).Methods("GET") // Rota para ver cadastros

	log.Fatal(http.ListenAndServe(":8080", router))
}
