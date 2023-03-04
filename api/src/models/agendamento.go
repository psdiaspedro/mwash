package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Agendamento struct {
	ID				uint64	`json:"id,omitempty"`
	PropriedadeID	uint64	`json:"propriedadeId,omitempty"`
	DiaAgendamento	string	`json:"diaAgendamento,omitempty"`
	Checkin			string	`json:"checkin,omitempty"`
	Checkout		string	`json:"checkout,omitempty"`
	Obs				string	`json:"obs,omitempty"`
}

type AgendamentoValor struct {
	Nome			string			`json:"nome,omitempty"`
	DiaAgendamento	string			`json:"diaAgendamento,omitempty"`
	Logadouro		string			`json:"logadouro,omitempty"`
	Valor			decimal.Decimal	`json:"valor,omitempty"`
}

func (agendamento *Agendamento) Preparar(etapa string) error {
	if erro := agendamento.Validar(etapa); erro != nil {
		return erro
	}

	agendamento.formatar()
	return nil
}

func (agendamento *Agendamento) Validar(etapa string) error {
	if etapa == "criando" {
		if agendamento.DiaAgendamento == "" || agendamento.Checkout == "" {
			return errors.New("esta faltando um ou mais dados obrigatórios para o agendamento: dia do agendamento e/ou checkout")
		}
	} else if etapa == "atualizando" {
		if agendamento.DiaAgendamento == "" && agendamento.Checkin == "" && agendamento.Checkout == "" && agendamento.Obs == "" {
			return errors.New("campos inválidos, você consegue atualizar um dos seguintes campos: data, checkin, checkout e/ou obs")
		}
	}

	return nil
}

func (agendamento *Agendamento) formatar() {
	agendamento.DiaAgendamento = strings.TrimSpace(agendamento.DiaAgendamento)
	agendamento.Checkin = strings.TrimSpace(agendamento.Checkin)
	agendamento.Checkout = strings.TrimSpace(agendamento.Checkout)
	agendamento.Obs = strings.TrimSpace(agendamento.Obs)
}

func (agendamento *Agendamento) GerarQueryString(agen Agendamento, agendamentoID uint64) (string, []any) {
	query := "update agendamentos set"
	var valores []any

	if agendamento.DiaAgendamento != "" {
		query += " dia_agendamento = STR_TO_DATE(?, '%d-%m-%Y')"
		valores = append(valores, agen.DiaAgendamento)
	}

	if agendamento.Checkin != "" {
		if agendamento.DiaAgendamento != "" {
			query += ","
		}
		query += " checkin = ?"
		valores = append(valores, agen.Checkin)
	} 

	if agendamento.Checkout != "" {
		if agendamento.Checkin != ""  || agendamento.DiaAgendamento != "" {
			query += ","
		}
		query += " checkout = ?"
		valores = append(valores, agen.Checkout)
	}

	if agendamento.Obs != "" {
		if agendamento.Checkout != "" || agendamento.Checkin != ""  || agendamento.DiaAgendamento != "" {
			query += ","
		}
		query += " observacoes = ?"
		valores = append(valores, agen.Obs)
	}
	
	query += " where id = ?"
	valores = append(valores, fmt.Sprintf("%d", agendamentoID))

	return query, valores
}