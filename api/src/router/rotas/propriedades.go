package rotas

import (
	"api/src/controllers"
)

var rotasPropriedades = []Route {
	{
		//Busca Todas as Propriedades do Usuário
		URI:		"/minhas_propriedades",
		Metodo:		"GET",
		Funcao:		controllers.ListarPropriedades,
		RequerAuth:	true,
	},
	{
		//Adiciona uma Propriedade
		URI:		"/minhas_propriedades",
		Metodo:		"POST",
		Funcao:		controllers.AdicionarPropriedade,
		RequerAuth:	true,
	},
	{
		//Atualiza uma Propriedade do usuario
		URI:		"/minhas_propriedades/{propriedadeId}",
		Metodo:		"PATCH",
		Funcao:		controllers.AtualizarPropriedade,
		RequerAuth:	true,
	},
	{
		//Remove uma Propriedade do usuario
		URI:		"/minhas_propriedades/{propriedadeId}",
		Metodo:		"DELETE",
		Funcao:		controllers.RemoverPropriedade,
		RequerAuth:	true,
	},
}  