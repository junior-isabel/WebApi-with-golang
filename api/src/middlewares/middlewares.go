package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"net/http"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Error(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
