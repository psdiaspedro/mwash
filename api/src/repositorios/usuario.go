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