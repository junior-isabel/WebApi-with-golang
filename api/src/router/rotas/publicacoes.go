package rotas

import (
	controller "api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
