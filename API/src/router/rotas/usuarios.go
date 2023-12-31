package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{ // SLICE DE ROTA CONTENDO TODAS AS ROTAS DE USUÁRIOS DA NOSSA API
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuariosId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
			URI:    "/usuarios/{usuariosId}",
			Metodo: http.MethodDelete,
			Funcao: controllers.DeletarUsuario,
			RequerAutenticacao: false,
	},
}