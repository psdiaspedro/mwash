package rotas

import (
	"api/src/controllers"
)

/*
Rotas relacionadas às informações do usuário
- Precisa estar logado
- Todos os tipos de usuário estão liberados
- Existe rotas de acesso do ADMIN e do Cliente
*/
var rotasUsuarios = []Route{
	{
		URI:        "/usuario",
		Metodo:     "GET",
		Funcao:     controllers.BuscarDadosUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuario/atualizar_dados",
		Metodo:     "PATCH",
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuario/senha",
		Metodo:     "PATCH",
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
