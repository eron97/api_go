package routes

/*

Esta pasta contém as definições das rotas da sua API.
Aqui você pode mapear os endpoints (URLs) da sua API para os handlers correspondentes.

*/

import (
	"net/http"

	"github.com/eron97/api_go/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Rota para a página de login
	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)

	// Rota para o dashboard do perfil
	router.HandleFunc("/dashboard", handlers.DashboardHandler).Methods(http.MethodGet)

	return router
}
