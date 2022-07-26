package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {

	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	lastId, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastId), nil

}

func (repositorio Usuarios) Buscar(nameOuNick string) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}
	return usuarios, nil
}

func (repositorio Usuarios) BuscarId(usuarioId uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioId)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {

	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	linha, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID)

	if erro != nil {
		return erro
	}

	_, erro = linha.LastInsertId()

	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(usuarioId uint64) error {

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID, &usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioIdFromToken uint64, usuarioId uint64) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into seguidores (usuario_id, seguidor_id) values(?,?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuarioId, usuarioIdFromToken)

	if erro != nil {
		return 0, erro
	}
	lastIdInsert, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInsert), nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioIdFromToken uint64, usuarioId uint64) (uint64, error) {

	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuarioId, usuarioIdFromToken)

	if erro != nil {
		return 0, erro
	}

	lastId, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastId), nil
}

func (repositorio Usuarios) BuscarSeguidor(usuarioId uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("select u.id, u.nome, u.nick from usuarios u inner join seguidores s on u.id = s.usuario_id where u.id = ?", usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarQuemSigo(usuarioIdFromToken uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("select u.id, u.nome, u.nick from usuarios u inner join seguidores s on s.usuario_id = u.id where s.seguidor_id = ?", usuarioIdFromToken)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
