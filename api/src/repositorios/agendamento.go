package repositorios

import (
	"api/src/models"
	"database/sql"
	// "errors"
	// "fmt"
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

func (repo Agendamento) BuscarAgendamentosPorData(data models.Data) ([]models.Agendamento, error) {
	query, valores := data.GerarQueryString(data)
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