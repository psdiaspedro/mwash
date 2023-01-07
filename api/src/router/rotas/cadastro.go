package rotas

import (
	"api/src/controllers"
)


/*
	Rota de cadastro de usuários da plataforma
	- Precisa estar logado
	- Apenas usuários ADMIN
*/
var rotaCadastro = Route {
	URI:		"/cadastrar",
	Metodo:		"POST",
	Funcao:		controllers.CadastrarUsuario,
	RequerAuth:	true,
}