package models

import (
	"errors"
	"fmt"
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
	Senha			string	`json:"senha,omitempty"`
	Acomodacao		string	`json:"acomodacao,omitempty"`
	Wifi			string	`json:"wifi,omitempty"`
	Outros			string	`json:"outros,omitempty"`
	Obs				string	`json:"obs,omitempty"`
	Cor				uint64	`json:"cor,omitempty"`
}

func (propriedade *Propriedade) GerarQueryString(prop Propriedade, propriedadeID uint64) (string, []any) {
	query := "update propriedades set"
	var valores []any

	if propriedade.Cidade != "" {
		query += " cidade = ?"
		valores = append(valores, prop.Cidade)
	}

	if propriedade.Bairro != "" {
		if propriedade.Cidade != "" {
			query += ","
		}
		query += " bairro = ?"
		valores = append(valores, prop.Bairro)
	} 

	if propriedade.CEP != "" {
		if propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " CEP = ?"
		valores = append(valores, prop.CEP)
	}

	if propriedade.Logadouro != "" {
		if propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " logadouro = ?"
		valores = append(valores, prop.Logadouro)
	}

	if propriedade.Numero != "" {
		if propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " numero = ?"
		valores = append(valores, prop.Numero)
	}

	if propriedade.Complemento != "" {
		if propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " complemento = ?"
		valores = append(valores, prop.Complemento)
	}

	if propriedade.Senha != "" {
		if propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}

		query += " senha = ?"
		valores = append(valores, prop.Cor)
	}

	if propriedade.Acomodacao != "" {
		if propriedade.Senha != "" || propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}

		query += " acomodacao = ?"
		valores = append(valores, prop.Cor)
	}
	
	if propriedade.Wifi != "" {
		if propriedade.Acomodacao != "" || propriedade.Senha != "" || propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " wifi = ?"
		valores = append(valores, prop.Obs)
	}

	if propriedade.Outros != "" {
		if propriedade.Wifi != "" || propriedade.Acomodacao != "" || propriedade.Senha != "" || propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " outros = ?"
		valores = append(valores, prop.Obs)
	}

	if propriedade.Obs != "" {
		if propriedade.Outros != "" || propriedade.Wifi != "" || propriedade.Acomodacao != "" || propriedade.Senha != "" || propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}
		query += " observacoes = ?"
		valores = append(valores, prop.Obs)
	}

	if propriedade.Cor > 0 {
		if propriedade.Obs != "" || propriedade.Outros != "" || propriedade.Wifi != "" || propriedade.Acomodacao != "" || propriedade.Senha != "" || propriedade.Complemento != "" || propriedade.Numero != "" || propriedade.Logadouro != "" || propriedade.CEP != "" || propriedade.Bairro != ""  || propriedade.Cidade != "" {
			query += ","
		}

		query += " cor = ?"
		valores = append(valores, prop.Cor)
	}

	valores = append(valores, fmt.Sprintf("%d", propriedadeID))
	query += " where id = ?"

	return query, valores
}

func (propriedade *Propriedade) Preparar(etapa string) error {
	if erro := propriedade.Validar(etapa); erro != nil {
		return erro
	}

	propriedade.formatar()
	return nil
}

func (propriedade *Propriedade) Validar(etapa string) error {
	if etapa == "cadastrar" {
		if propriedade.Cidade == "" || propriedade.Bairro == "" ||   propriedade.CEP == "" ||  propriedade.Logadouro == "" ||  propriedade.Numero == "" {
			return errors.New("esta faltando um ou mais dados obrigatórios para o cadastro (cidade, bairro, CEP, logadouro ou número)")
		}
	} else if etapa == "atualizar" {
		if propriedade.Cidade == "" && propriedade.Bairro == "" &&   propriedade.CEP == "" &&  propriedade.Logadouro == "" &&  propriedade.Numero == "" &&  propriedade.Complemento == ""  && propriedade.Senha == "" && propriedade.Acomodacao == "" && propriedade.Wifi == "" && propriedade.Obs == "" && propriedade.Outros == "" {
			return errors.New("campo invalido, você consegue atualizar um dos seguintes campos: cidade, bairro, cep, logadouro, numero, complemento, senha, acomodação, wifi, outros, observações e senha")
		}
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
	propriedade.Senha = strings.TrimSpace(propriedade.Obs)
	propriedade.Acomodacao = strings.TrimSpace(propriedade.Obs)
	propriedade.Wifi = strings.TrimSpace(propriedade.Obs)
	propriedade.Outros = strings.TrimSpace(propriedade.Obs)
	propriedade.Obs = strings.TrimSpace(propriedade.Obs)
}

func (propriedade *Propriedade) PropriedadeCadastrada() error {
	if propriedade.ID == 0 {
		return errors.New("propriedade informada não encontrada")
	}

	return nil
}