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
		RequerAuth:	false,
	},
	{
		//Adiciona uma Propriedade
		URI:		"/minhas_propriedades",
		Metodo:		"POST",
		Funcao:		controllers.AdicionarPropriedade,
		RequerAuth:	false,
	},
	{
		//Atualiza uma Propriedade
		URI:		"/minhas_propriedades",
		Metodo:		"PUT",
		Funcao:		controllers.AtualizarPropriedade,
		RequerAuth:	false,
	},
	{
		//Remove uma Propriedade
		URI:		"/minhas_propriedades",
		Metodo:		"DELETE",
		Funcao:		controllers.RemoverPropriedade,
		RequerAuth:	false,
	},
}  