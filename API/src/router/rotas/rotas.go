package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota REPRESENTA TODAS AS ROTAS DA API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar COLOCA TODAS AS ROTAS DENTRO DO ROUTER
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios	// ROTAS CRIADAS NO ARQUIVO USUÁRIOS

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}