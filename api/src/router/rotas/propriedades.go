package rotas

import (
	"api/src/controllers"
)

/*
	Rotas relacionadas às propriedades do cliente
	- Precisa estar logado
	- Todos os tipos de usuário estão liberados
*/
var rotasPropriedades = []Route {
	{
		//Busca todas as propriedades do usuário logado
		URI:		"/minhas_propriedades",
		Metodo:		"GET",
		Funcao:		controllers.ListarPropriedades,
		RequerAuth:	true,
	},
	{
		//Adiciona uma propriedade para o usuário logado
		URI:		"/minhas_propriedades",
		Metodo:		"POST",
		Funcao:		controllers.AdicionarPropriedade,
		RequerAuth:	true,
	},
	{
		//Atualiza uma propriedade do usuário logado
		URI:		"/minhas_propriedades/{propriedadeId}",
		Metodo:		"PATCH",
		Funcao:		controllers.AtualizarPropriedade,
		RequerAuth:	true,
	},
	{
		//Remove uma propriedade do usuário logado
		URI:		"/minhas_propriedades/{propriedadeId}",
		Metodo:		"DELETE",
		Funcao:		controllers.RemoverPropriedade,
		RequerAuth:	true,
	},
}  