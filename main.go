package main

/*

Este é o arquivo principal do seu projeto.
Aqui você iniciará o servidor HTTP e configurará as dependências necessárias para sua API.

*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eron97/api_go/routes"
)

func main() {
	// Configuração das rotas
	router := routes.SetupRouter()

	// Iniciando o servidor HTTP
	port := ":8080" // Porta em que o servidor será executado
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
