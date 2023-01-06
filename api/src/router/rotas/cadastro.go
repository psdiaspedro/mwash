package rotas

import (
	"api/src/controllers"
)


/*
	Rota de cadastro de usuários da plataforma
	- Possuí apenas /cadastrar
	- Precisa estar aut
	- Chama a função CadastrarUsuao
*/
var rotaCadastro = Route {
	URI:		"/cadastrar",
	Metodo:		"POST",
	Funcao:		controllers.CadastrarUsuario,
	RequerAuth:	false,
}