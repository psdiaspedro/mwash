package models

import "errors"

type Senha struct {
	Nova	string `json:"nova"`
	Atual	string `json:"atual"`
}

func (senha *Senha) Preparar() error {
	if erro := senha.Validar(); erro != nil {
		return erro
	}

	return nil
}

func (senha *Senha) Validar() error {
	if senha.Nova == "" || senha.Atual == "" {
		return errors.New("campos invalidos ou valores vazios, por favor tente novamente")
	}

	return nil
}
