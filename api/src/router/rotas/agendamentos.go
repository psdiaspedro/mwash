package rotas

import (
	"api/src/controllers"
)

var rotasAgendamentos = []Route {
	{
		//Busca Todos os Agendamentos do Usuario
		URI:		"/meus_agendamentos",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamento,
		RequerAuth:	true,
	}, 
	{
		//Adiciona um Agendamento
		URI:		"/agendar",
		Metodo:		"POST",  
		Funcao:		controllers.AdicionarAgendamento,
		RequerAuth:	true,
	},
	{
		//Atualiza um Agendamento 
		URI:		"/meus_agendamentos",
		Metodo:		"PUT", 
		Funcao:		controllers.AtualizarAgendamento,
		RequerAuth:	true, 
	},
	{
		//Remove um Agendamento
		URI:		"/meus_agendamentos",
		Metodo:		"DELETE",  
		Funcao:		controllers.RemoverAgendamento,
		RequerAuth:	true,
	},
}  