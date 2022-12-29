package models

import (
	"errors"
	"strings"
)

type Agendamento struct {
	ID				uint64	`json:"id,omitempty"`
	PropriedadeID	uint64	`json:"propriedadeId,omitempty"`
	DiaAgendamento	string	`json:"diaAgendamento,omitempty"`
	Checkin			string	`json:"checkin,omitempty"`
	Checkout		string	`json:"checkout,omitempty"`
	Obs				string	`json:"obs,omitempty"`
}

func (agendamento *Agendamento) Preparar() error {
	if erro := agendamento.Validar(); erro != nil {
		return erro
	}

	agendamento.formatar()
	return nil
}

func (agendamento *Agendamento) Validar() error {
	if agendamento.DiaAgendamento == "" || agendamento.Checkout == "" {
		return errors.New("campos obrigatórios: dia do agendamento e checkout")
	}

	return nil
}

func (agendamento *Agendamento) formatar() {
	agendamento.DiaAgendamento = strings.TrimSpace(agendamento.DiaAgendamento)
	agendamento.Checkin = strings.TrimSpace(agendamento.Checkin)
	agendamento.Checkout = strings.TrimSpace(agendamento.Checkout)
	agendamento.Obs = strings.TrimSpace(agendamento.Obs)
}