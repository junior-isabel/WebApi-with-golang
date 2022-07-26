package controller

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro := json.Unmarshal(body, &usuario); erro != nil {

		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar("cadastro"); erro != nil {
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
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {

		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()

	if erro != nil {

		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.Buscar("")

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario, erro := repositorio.BuscarId(usuarioID)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametro := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIdFromToken, erro := autenticacao.ExtrairUsuarioIDToken(r)

	if erro != nil {
		respostas.Error(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioId != usuarioIdFromToken {
		respostas.Error(w, http.StatusForbidden, errors.New("não pode actualizar dados de outros utilizadores, precisa de permissão"))
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario

	if erro := json.Unmarshal(bodyRequest, &usuario); erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	if erro := usuario.Preparar("edicao"); erro != nil {
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
	erro = repositorio.Atualizar(usuarioId, usuario)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
	}

	w.Write([]byte(fmt.Sprintf("Atualizar um usuario %d", usuarioId)))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("eliminar um usuario"))
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioId == 0 {
		respostas.Error(w, http.StatusBadRequest, errors.New("usuario invalido"))
		return
	}

	usuarioIdFromToken, erro := autenticacao.ExtrairUsuarioIDToken(r)

	if erro != nil {
		respostas.Error(w, http.StatusForbidden, erro)
		return
	}

	if usuarioId == usuarioIdFromToken {
		respostas.Error(w, http.StatusConflict, errors.New("usuario não pode se seguir"))
		return
	}
	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if _, erro := repositorio.Seguir(usuarioIdFromToken, usuarioId); erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, nil)
}

func BuscarSeguidor(w http.ResponseWriter, r *http.Request) {

	usuarioId, erro := autenticacao.ExtrairUsuarioIDToken(r)
	if erro != nil {
		respostas.Error(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.BuscarSeguidor(usuarioId)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarQuemSigo(w http.ResponseWriter, r *http.Request) {

	usuarioId, erro := autenticacao.ExtrairUsuarioIDToken(r)

	if erro != nil {
		respostas.Error(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.BuscarQuemSigo(usuarioId)

	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}
