package rotas

import (
	controller "api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controller.Login,
	RequerAutenticacao: false,
}
