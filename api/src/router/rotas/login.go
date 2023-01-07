package rotas

import (
	"api/src/controllers"
)

/*
	Rota principal para acesso à plataforma
	- Não precisa estar logado
	- Todos os tipos de usuário estão liberados
*/
var rotaLogin = Route {
	URI:		"/login",
	Metodo:		"POST",
	Funcao:		controllers.Login,
	RequerAuth:	false,
} 