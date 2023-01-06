package rotas

import (
	"api/src/controllers"
)

/*
	Rota principal para acesso à plataforma
	- Possui apenas /login
	- Não precisa de auth pois é a porta de entrada
	- Chama a função Login
*/
var rotaLogin = Route {
	URI:		"/login",
	Metodo:		"POST",
	Funcao:		controllers.Login,
	RequerAuth:	false,
} 