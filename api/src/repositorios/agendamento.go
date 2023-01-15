package repositorios

import (
	"api/src/models"
	"database/sql"
	"strconv"
)

type Agendamento struct {
	db *sql.DB
}

func NovoRepoAgendamento(db *sql.DB) *Agendamento {
	return &Agendamento{db}
}

func (repo Agendamento) CriarAgendamento(agendamento models.Agendamento) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into agendamentos (propriedade_id, dia_agendamento, checkin, checkout, observacoes) values (?, STR_TO_DATE(?, '%d-%m-%Y'), ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(agendamento.PropriedadeID, agendamento.DiaAgendamento, agendamento.Checkin, agendamento.Checkout, agendamento.Obs)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repo Agendamento) BuscarAgendamentosDoUsuario(usuarioID uint64) ([]models.Agendamento, error) {
	linhas, erro := repo.db.Query("select a.id, a.propriedade_id, a.dia_agendamento, a.checkin, a.checkout, a.observacoes from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id inner join usuarios u on u.id = p.cliente_id where u.id = ? order by a.dia_agendamento", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var agendamentos []models.Agendamento

	for linhas.Next() {
		var agendamento models.Agendamento

		if erro = linhas.Scan(
			&agendamento.ID,
			&agendamento.PropriedadeID,
			&agendamento.DiaAgendamento,
			&agendamento.Checkin,
			&agendamento.Checkout,
			&agendamento.Obs,
		); erro != nil {
			return nil, erro
		}

		agendamentos = append(agendamentos, agendamento)
	}

	return agendamentos, nil
}

func (repo Agendamento) BuscarAgendamentosPropriedade(propriedadeID uint64) ([]models.Agendamento, error) {
	linhas, erro := repo.db.Query("select a.id, a.propriedade_id, a.dia_agendamento, a.checkin, a.checkout, a.observacoes from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id where p.id = ?", propriedadeID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var agendamentos []models.Agendamento

	for linhas.Next() {
		var agendamento models.Agendamento

		if erro = linhas.Scan(
			&agendamento.ID,
			&agendamento.PropriedadeID,
			&agendamento.DiaAgendamento,
			&agendamento.Checkin,
			&agendamento.Checkout,
			&agendamento.Obs,
		); erro != nil {
			return nil, erro
		}

		agendamentos = append(agendamentos, agendamento)
	}

	return agendamentos, nil
}

func (repo Agendamento) BuscarAgendamentosPorDataLogado(data models.Data, usuarioId uint64) ([]models.Agendamento, error) {
	query, valores := data.GerarQueryStringUsuarioId(data, usuarioId)
	linhas, erro := repo.db.Query(query, valores...)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var agendamentos []models.Agendamento

	for linhas.Next() {
		var agendamento models.Agendamento

		if erro = linhas.Scan(
			&agendamento.ID,
			&agendamento.PropriedadeID,
			&agendamento.DiaAgendamento,
			&agendamento.Checkin,
			&agendamento.Checkout,
			&agendamento.Obs,
		); erro != nil {
			return nil, erro
		}

		agendamentos = append(agendamentos, agendamento)
	}

	return agendamentos, nil
}

func (repo Agendamento) BuscarAgendamentoPorId(agendamentoID uint64) (models.Agendamento, error) {
	linha, erro := repo.db.Query("select * from agendamentos where id = ?", agendamentoID)
	if erro != nil {
		return models.Agendamento{}, erro
	}
	defer linha.Close()

	var agendamento models.Agendamento

	if linha.Next() {
		if erro = linha.Scan(
			&agendamento.ID,
			&agendamento.PropriedadeID,
			&agendamento.DiaAgendamento,
			&agendamento.Checkin,
			&agendamento.Checkout,
			&agendamento.Obs,
		); erro != nil {
			return models.Agendamento{}, erro
		}
	}

	return agendamento, nil
}

func (repo Agendamento) AtualizarAgendamento(agendamentoID uint64, agendamento models.Agendamento) error {
	query, valores := agendamento.GerarQueryString(agendamento, agendamentoID)
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

func (repo Agendamento) DeletarAgendamento(agendamentoID uint64) error {
	statement, erro := repo.db.Prepare("delete from agendamentos where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(agendamentoID); erro != nil {
		return erro
	}

	return nil
}

func (repo Agendamento) BuscarClientePorAgendamentoId(agendamentoID uint64) (uint64, error) {
	linha, erro := repo.db.Query("select p.cliente_id from propriedades p INNER JOIN agendamentos a ON p.id = a.propriedade_id where a.id = ?", agendamentoID)
	if erro != nil {
		return 0, erro
	}
	defer linha.Close()

	var clienteIDString string

	if linha.Next() {
		if erro = linha.Scan(&clienteIDString); erro != nil {
			return 0, erro
		}
	}

	var clienteID uint64
	if clienteIDString == "" {
		clienteID = 0
	} else {
		clienteID, erro = strconv.ParseUint(clienteIDString, 10, 64)
		if erro != nil {
			return 0, erro
		}
	}
	return clienteID, nil
}
