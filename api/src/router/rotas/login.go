package rotas

import (
	"api/src/controllers"
)

//Envia os dados do login
var rotaLogin = Route {
	URI:		"/login",
	Metodo:		"POST",
	Funcao:		controllers.Login,
	RequerAuth:	false,
} 