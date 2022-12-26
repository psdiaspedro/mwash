package rotas

import (
	"api/src/controllers"
)

var rotasAgendamentos = []Route {
	{
		//Busca Todos os Agendamentos do Usuario
		URI:		"/agendamentos",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentos,
		RequerAuth:	true,
	},
	{
		//Busca Todos os Agendamentos de uma propriedade
		URI:		"/agendamentos/propriedades/{propriedadeId}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPropriedade,
		RequerAuth:	true,
	},
	{
		//Adiciona um Agendamento
		URI:		"/agendamentos/propriedades/{propriedadeId}",
		Metodo:		"POST",  
		Funcao:		controllers.AdicionarAgendamento,
		RequerAuth:	true,
	},
	{
		//Atualiza um Agendamento 
		URI:		"/agendamentos",
		Metodo:		"PATCH", 
		Funcao:		controllers.AtualizarAgendamento,
		RequerAuth:	true, 
	},
	{
		//Remove um Agendamento
		URI:		"/agendamentos",
		Metodo:		"DELETE",  
		Funcao:		controllers.RemoverAgendamento,
		RequerAuth:	true,
	},
}  