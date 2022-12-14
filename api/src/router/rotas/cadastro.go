package rotas

import (
	"api/src/controllers"
)

//Envia os dados para cadastro
var rotaCadastro = Route {
	URI:		"/cadastrar",
	Metodo:		"POST",
	Funcao:		controllers.CadastrarUsuario,
	RequerAuth:	false,
}