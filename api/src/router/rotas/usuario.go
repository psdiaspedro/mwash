package rotas

import (
	"api/src/controllers"
)

var rotasUsuarios = []Route {
	{
		//Pega Todas Informações do Usuario logado
		URI:		"/usuario",
		Metodo:		"GET",
		Funcao:		controllers.BuscarDadosUsuario, 
		RequerAuth:	true,
	},
	{
		//Atualiza os dados do usuário
		URI:		"/usuario/atualizar_dados",
		Metodo:		"PATCH",
		Funcao:		controllers.AtualizarUsuario,
		RequerAuth:	true,
	},
	{
		//Atualiza a senha do usuário
		URI:		"/usuario/senha",
		Metodo:		"PATCH",
		Funcao:		controllers.AtualizarSenha,
		RequerAuth:	true,
	}, 
} 