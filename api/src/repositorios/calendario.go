package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Calendario struct {
	db *sql.DB
}

func NovoRepoCalendario(db *sql.DB) *Calendario {
	return &Calendario{db}
}

func (repo Calendario) BuscarAgendamentosPorData(data models.Data) ([]models.Calendario, error) {
	query, valores := data.GerarQueryString(data)
	linhas, erro := repo.db.Query(query, valores...)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var calendarios []models.Calendario

	for linhas.Next() {
		var calendario models.Calendario

		if erro = linhas.Scan(
			&calendario.AgendamentoID,
			&calendario.PropriedadeID,
			&calendario.DiaAgendamento,
			&calendario.Checkin,
			&calendario.Checkout,
			&calendario.Obs,
			&calendario.ProprietarioID,
			&calendario.Cidade,
			&calendario.Bairro,
			&calendario.CEP,
			&calendario.Logadouro,
			&calendario.Numero,
			&calendario.Complemento,
			&calendario.Nome,
			&calendario.Email,
			&calendario.Contato,
		); erro != nil {
			return nil, erro
		}

		calendarios = append(calendarios, calendario)
	}

	return calendarios, nil
}

func (repo Calendario) BuscarAgendamentosDoUsuario(usuarioID uint64) ([]models.Calendario, error) {
	linhas, erro := repo.db.Query("select a.*, p.cliente_id, p.cidade, p.bairro, p.CEP, p.logadouro, p.numero, p.complemento, u.nome, u.email, u.contato from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id INNER JOIN usuarios u on u.id = p.cliente_id where u.id = ? order by a.dia_agendamento desc", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var agendamentos []models.Calendario

	for linhas.Next() {
		var agendamento models.Calendario

		if erro = linhas.Scan(
			&agendamento.AgendamentoID,
			&agendamento.PropriedadeID,
			&agendamento.DiaAgendamento,
			&agendamento.Checkin,
			&agendamento.Checkout,
			&agendamento.Obs,
			&agendamento.ProprietarioID,
			&agendamento.Cidade,
			&agendamento.Bairro,
			&agendamento.CEP,
			&agendamento.Logadouro,
			&agendamento.Numero,
			&agendamento.Complemento,
			&agendamento.Nome,
			&agendamento.Email,
			&agendamento.Contato,
		); erro != nil {
			return nil, erro
		}

		agendamentos = append(agendamentos, agendamento)
	}

	return agendamentos, nil
}
