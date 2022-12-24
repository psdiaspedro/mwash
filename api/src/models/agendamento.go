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

// func (propriedade *Propriedade) GerarQueryString(prop Propriedade, propriedadeID uint64) (string, []any) {
// 	query := "update propriedades set"
// 	var valores []any

// 	if propriedade.Cidade != "" {
// 		query += " cidade = ?"
// 		valores = append(valores, prop.Cidade)
// 	}

// 	if propriedade.Bairro != "" {
// 		if propriedade.Cidade != "" {
// 			query += ","
// 		}
// 		query += " bairro = ?"
// 		valores = append(valores, prop.Bairro)
// 	} 

// 	if propriedade.CEP != "" {
// 		if propriedade.Bairro != ""  || propriedade.Cidade != "" {
// 			query += ","
// 		}
// 		query += " CEP = ?"
// 		valores = append(valores, prop.CEP)
// 	}

// 	if propriedade.Logadouro != "" {
// 		if propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
// 			query += ","
// 		}
// 		query += " logadouro = ?"
// 		valores = append(valores, prop.Logadouro)
// 	}

// 	if propriedade.Numero != "" {
// 		if propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
// 			query += ","
// 		}
// 		query += " numero = ?"
// 		valores = append(valores, prop.Numero)
// 	}

// 	if propriedade.Complemento != "" {
// 		if propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
// 			query += ","
// 		}
// 		query += " complemento = ?"
// 		valores = append(valores, prop.Complemento)
// 	}
	
// 	valores = append(valores, fmt.Sprintf("%d", propriedadeID))
// 	query += " where id = ?"

// 	return query, valores
// }

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