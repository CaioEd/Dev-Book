package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar VAI RETORNAR UM ROUTER COM AS ROTAS CONFIGURADAS
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}