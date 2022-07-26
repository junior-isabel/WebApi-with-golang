package rotas

import (
	controller "api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeleteUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/buscar/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarSeguidor,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/buscar/quem-segue",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarQuemSigo,
		RequerAutenticacao: true,
	},
}
