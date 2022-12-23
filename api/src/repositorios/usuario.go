package repositorios

import (
	"api/src/models"
	"database/sql"
)

type usuario struct {
	db	*sql.DB
}

func NovoRepoUsuario(db *sql.DB) *usuario {
	return &usuario{db}
}

func (repo usuario) CadastrarUsuarioNoBanco(usuario models.Usuario) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into usuarios (nome, email, senha, contato, admin) values (?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Email, usuario.Senha, usuario.Contato, usuario.Admin)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repo usuario) BuscarUsuarioPorEmail(email string) (models.Usuario, error) {
	linha, erro := repo.db.Query("select id, nome, senha, admin from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Senha, &usuario.Admin); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repo usuario) BuscarUsuarioPorId(usuarioId uint64) (models.Usuario, error) {
	linha, erro := repo.db.Query("select id, nome, email, contato, admin from usuarios where id = ?", usuarioId)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Contato,
			&usuario.Admin,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repo usuario) AtualizarDadosUsuario(usuarioID uint64, usuario models.Usuario) error {
	query, valores := usuario.GerarQueryString(usuario, usuarioID)
	statement, erro := repo.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(valores...); erro != nil {
		return erro
	}

	return nil

}

func (repo usuario) BuscarSenhaAtualUsuario(usuarioID uint64) (string, error) {
	linha, erro := repo.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repo usuario) AtualizarSenha(usuarioID uint64, senha string) error{
	statement, erro := repo.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}