package rotas

import (
	"api/src/controllers"
)

var rotasUsuarios = []Route {
	{
		//Pega Todas Informações do Usuario
		URI:		"/usuario",
		Metodo:		"GET",
		Funcao:		controllers.BuscarUsuario, 
		RequerAuth:	true,
	},
	{
		//Atualiza os dados do usuário
		URI:		"/usuario/atualizar_dados",
		Metodo:		"PUT",
		Funcao:		controllers.AtualizarUsuario,
		RequerAuth:	true,
	},
	{
		//Atualiza a senha do usuário
		URI:		"/usuario/senha",
		Metodo:		"PUT",
		Funcao:		controllers.AtualizarSenha,
		RequerAuth:	true,
	}, 
} 