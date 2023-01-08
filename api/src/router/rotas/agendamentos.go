package rotas

import (
	"api/src/controllers"
)

/*
	Rotas relacionadas àos agendamentos do cliente
	- Precisa estar logado
	- Todos os tipos de usuário estão liberados
	- Existe rotas de acesso do ADMIN e do Cliente
*/
var rotasAgendamentos = []Route {
	{
		URI:		"/agendamentos",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentos,
		RequerAuth:	true,
	},
	{
		URI:		"/agendamentos/propriedades/{propriedadeId}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPropriedade,
		RequerAuth:	true,
	},
	{
		URI:		"/agendamentos/{data}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPorData,
		RequerAuth:	true,
	},
	{
		URI:		"/agendamentos/usuario/{data}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosPorDataLogado,
		RequerAuth:	true,
	},
	{
		URI:		"/agendamentos/adm/{usuarioId}",
		Metodo:		"GET",
		Funcao:		controllers.BuscarAgendamentosUsuarioId,
		RequerAuth:	true,
	},
	{
		URI:		"/agendamentos/propriedades/{propriedadeId}",
		Metodo:		"POST",  
		Funcao:		controllers.AdicionarAgendamento,
		RequerAuth:	true,
	},
	{
		// Cliente - Atualiza um Agendamento especifico
		URI:		"/agendamentos/{agendamentoId}",
		Metodo:		"PATCH", 
		Funcao:		controllers.AtualizarAgendamento,
		RequerAuth:	true, 
	},
	{
		// Cliente - Remove um Agendamento
		URI:		"/agendamentos/{agendamentoId}",
		Metodo:		"DELETE",  
		Funcao:		controllers.RemoverAgendamento,
		RequerAuth:	true,
	},
}  