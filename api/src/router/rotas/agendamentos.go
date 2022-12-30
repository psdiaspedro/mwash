package rotas

import (
	"api/src/controllers"
)

var rotasAgendamentos = []Route {
	{
		//Busca Todos os Agendamentos do Usuario - 
		URI:		"/agendamentos",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentos,
		RequerAuth:	true,
	},
	{
		//Busca Todos os Agendamentos de uma propriedade - Admin
		URI:		"/agendamentos/propriedades/{propriedadeId}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPropriedade,
		RequerAuth:	true,
	},
	{
		//Busca Todos os Agendamentos baseado em uma data - Admin
		URI:		"/agendamentos/{data}", //data -> AAAA/MM/DD - 2006/01/02
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPorData,
		RequerAuth:	true,
	},
	{
		//Busca Todos os Agendamentos do ususario logado baseado em uma data
		URI:		"/agendamentos/usuario/{data}", //data -> AAAA/MM/DD - 2006/01/02
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPorDataLogado,
		RequerAuth:	true,
	},
	{
		//Busca Todos os Agendamentos de um usuario especificado - Admin
		URI:		"/agendamentos/{usuarioId}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosUsuarioId,
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
		URI:		"/agendamentos/{agendamentoId}",
		Metodo:		"PATCH", 
		Funcao:		controllers.AtualizarAgendamento,
		RequerAuth:	true, 
	},
	{
		//Remove um Agendamento
		URI:		"/agendamentos/{agendamentoId}",
		Metodo:		"DELETE",  
		Funcao:		controllers.RemoverAgendamento,
		RequerAuth:	true,
	},
}  