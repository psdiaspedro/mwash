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
		"insert into propriedades (cliente_id, cidade, bairro, CEP, logadouro, numero, complemento, senha, acomodacao, wifi, outros, observacoes, cor, valor) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(propriedade.ProprietarioID, propriedade.Cidade, propriedade.Bairro, propriedade.CEP, propriedade.Logadouro, propriedade.Numero, propriedade.Complemento, propriedade.Senha, propriedade.Acomodacao, propriedade.Wifi, propriedade.Outros, propriedade.Obs, propriedade.Cor, propriedade.Valor)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repo Propriedade) BuscarPropriedadesDoUsuario(usuarioID uint64) ([]models.Propriedade, error) {
	linhas, erro := repo.db.Query("select p.id, p.cidade, p.bairro, p.CEP, p.logadouro, p.numero, p.complemento, p.senha, p.acomodacao, p.wifi, p.outros, p.observacoes, p.cor, p.valor from propriedades p where cliente_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var propriedades []models.Propriedade

	for linhas.Next() {
		var propriedade models.Propriedade

		if erro = linhas.Scan(
			&propriedade.ID,
			&propriedade.Cidade,
			&propriedade.Bairro,
			&propriedade.CEP,
			&propriedade.Logadouro,
			&propriedade.Numero,
			&propriedade.Complemento,
			&propriedade.Senha,
			&propriedade.Acomodacao,
			&propriedade.Wifi,
			&propriedade.Outros,
			&propriedade.Obs,
			&propriedade.Cor,
			&propriedade.Valor,
		); erro != nil {
			return nil, erro
		}

		propriedades = append(propriedades, propriedade)
	}

	return propriedades, nil
}

func (repo Propriedade) BuscarTodasPropriedades() ([]models.Propriedade, error) {
	linhas, erro := repo.db.Query("select p.id, p.cidade, p.bairro, p.CEP, p.logadouro, p.numero, p.complemento, p.senha, p.acomodacao, p.wifi, p.outros, p.observacoes, p.cor, p.valor from propriedades p")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var propriedades []models.Propriedade

	for linhas.Next() {
		var propriedade models.Propriedade

		if erro = linhas.Scan(
			&propriedade.ID,
			&propriedade.Cidade,
			&propriedade.Bairro,
			&propriedade.CEP,
			&propriedade.Logadouro,
			&propriedade.Numero,
			&propriedade.Complemento,
			&propriedade.Senha,
			&propriedade.Acomodacao,
			&propriedade.Wifi,
			&propriedade.Outros,
			&propriedade.Obs,
			&propriedade.Cor,
			&propriedade.Valor,
		); erro != nil {
			return nil, erro
		}

		propriedades = append(propriedades, propriedade)
	}

	return propriedades, nil
}

func (repo Propriedade) BuscarPropriedadePorId(propriedadeID uint64) (models.Propriedade, error) {
	linha, erro := repo.db.Query("select * from propriedades where id = ?", propriedadeID)
	if erro != nil {
		return models.Propriedade{}, erro
	}
	defer linha.Close()

	var propriedade models.Propriedade

	if linha.Next() {
		if erro = linha.Scan(
			&propriedade.ID,
			&propriedade.ProprietarioID,
			&propriedade.Cidade,
			&propriedade.Bairro,
			&propriedade.CEP,
			&propriedade.Logadouro,
			&propriedade.Numero,
			&propriedade.Complemento,
			&propriedade.Senha,
			&propriedade.Acomodacao,
			&propriedade.Wifi,
			&propriedade.Outros,
			&propriedade.Obs,
			&propriedade.Cor,
			&propriedade.Valor,
		); erro != nil {
			return models.Propriedade{}, erro
		}
	}

	return propriedade, nil
}

func (repo Propriedade) AtualizarPropriedade(propriedadeID uint64, propriedade models.Propriedade) error {
	query, valores := propriedade.GerarQueryString(propriedade, propriedadeID)
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

func (repo Propriedade) DeletarPropriede(propriedadeID uint64) error {
	statement, erro := repo.db.Prepare("delete from propriedades where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(propriedadeID); erro != nil {
		return erro
	}

	return nil
}
