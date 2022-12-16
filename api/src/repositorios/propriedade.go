package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Propriedade struct {
	db *sql.DB
}

func NovoRepoPropriedade(db *sql.DB) *Propriedade {
	return &Propriedade{db}
}

func (repo Propriedade) CriarPropriedade(propriedade models.Propriedade) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into propriedades (cliente_id, cidade, bairro, CEP, logadouro, numero, complemento) values (?, ?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(propriedade.ProprietarioID, propriedade.Cidade, propriedade.Bairro, propriedade.CEP, propriedade.Logadouro, propriedade.Numero, propriedade.Complemento)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}