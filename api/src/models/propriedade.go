package models

import (
	"errors"
	"strings"
)

type Propriedade struct {
	ID				uint64	`json:"id,omitempty"`
	ProprietarioID	uint64	`json:"proprietarioid,omitempty"`
	Cidade			string	`json:"cidade,omitempty"`
	Bairro			string	`json:"bairro,omitempty"`
	CEP				string	`json:"cep,omitempty"`
	Logadouro		string	`json:"logadouro,omitempty"`
	Numero			string	`json:"numero,omitempty"`
	Complemento		string	`json:"complemento,omitempty"`
}

func (propriedade *Propriedade) Preparar() error {
	if erro := propriedade.Validar(); erro != nil {
		return erro
	}

	propriedade.formatar()
	return nil
}

func (propriedade *Propriedade) Validar() error {
	if propriedade.Cidade == "" {
		return errors.New("cidade é obrigatório")
	}

	if propriedade.Bairro == "" {
		return errors.New("bairro é obrigatório")
	}

	if propriedade.CEP == "" {
		return errors.New("CEP é obrigatório")
	}

	if propriedade.Logadouro == "" {
		return errors.New("Logadouro é obrigatório")
	}

	if propriedade.Numero == "" {
		return errors.New("Numero é obrigatório")
	}

	return nil
}

func (propriedade *Propriedade) formatar() {
	propriedade.Cidade = strings.TrimSpace(propriedade.Cidade)
	propriedade.Bairro = strings.TrimSpace(propriedade.Bairro)
	propriedade.CEP = strings.TrimSpace(propriedade.CEP)
	propriedade.Logadouro = strings.TrimSpace(propriedade.Logadouro)
	propriedade.Numero = strings.TrimSpace(propriedade.Numero)
	propriedade.Complemento = strings.TrimSpace(propriedade.Complemento)
}