package controller

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioSalvoNobanco, erro := repositorio.BuscarPorEmail(usuario.Email)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerifcarSenha(usuarioSalvoNobanco.Senha, usuario.Senha); erro != nil {
		respostas.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNobanco.ID)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))

}
