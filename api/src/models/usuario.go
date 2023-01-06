package models

import (
	"api/src/seguranca"
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID			uint64	`json:"id,omitempty"`
	Nome		string	`json:"nome,omitempty"`
	Email		string	`json:"email,omitempty"`
	Senha		string	`json:"senha,omitempty"`
	Contato		string	`json:"contato,omitempty"`
	Admin		bool	`json:"admin,omitempty"`
	Token		string	`json:"token,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if etapa == "cadastro" {
		if usuario.Nome == "" || usuario.Email == "" || usuario.Senha == "" || 
		usuario.Contato == "" {
			return errors.New("esta faltando um ou mais dados obrigatórios para o cadastro (nome, email, senha ou contato)")
		}
	
		if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
			return errors.New("email no formato invalido")
		}
	} else if etapa == "atualizar" {
		if usuario.Nome == "" && usuario.Email == "" && usuario.Contato == "" {
			return errors.New("campos inválidos, você consegue atualizar um dos seguintes campos: nome, email e/ou contato")
		}

		if usuario.Senha != "" {
			return errors.New("ops, esse não é o lugar certo para atualizar sua senha")
		}
		
		if usuario.Email != "" {
			if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
				return errors.New("email no formato invalido")
			}
		}
	} else if etapa == "login" {
		if usuario.Email == "" || usuario.Senha == "" {
			return errors.New("campos obrigatórios para o login: email e senha")
		}

		if usuario.Email != "" {
			if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
				return errors.New("email no formato invalido")
			}
		}
	}
	
	return nil
}

func (usuario *Usuario) GerarQueryString(user Usuario, usuarioID uint64) (string, []any) {
	query := "update usuarios set"
	var valores []any

	if usuario.Nome != "" {
		query += " nome = ?"
		valores = append(valores, user.Nome)
	}

	if usuario.Email != "" {
		if usuario.Nome != "" {
			query += ","
		}
		query += " email = ?"
		valores = append(valores, user.Email)
	} 

	if usuario.Contato != "" {
		if usuario.Email != ""  || usuario.Nome != "" {
			query += ","
		}
		query += " Contato = ?"
		valores = append(valores, user.Contato)
	}

	valores = append(valores, fmt.Sprintf("%d", usuarioID))
	query += " where id = ?"

	return query, valores
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, erro := seguranca.GerarHash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaHash)
	}

	return nil
}

func (usuario *Usuario) UsuarioCadastrado() error {
	if usuario.ID == 0 {
		return errors.New("email e/ou senha informado não encontrou um usuário válido")
	}

	return nil
}