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
		); erro != nil {
			return nil, erro
		}

		calendarios = append(calendarios, calendario)
	}

	return calendarios, nil
}
