package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"

	//"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
Função chamada pela rota GET /agendamentos
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, impede o acesso
  - Recupera o ID do usário logado através do token
  - Chama a função que busca todos os agendamentos do usuário logado
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 204
  - lista de agendamentos do usuário logado

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarAgendamentos(w http.ResponseWriter, r *http.Request) {
	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("eu sei que você é admin e pode fazer tudo, mas essa rota é exclusiva do cliente"))
		return
	}

	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var agendamentos []models.Calendario

	repo := repositorios.NovoRepoCalendario(db)
	agendamentos, erro = repo.BuscarAgendamentosDoUsuario(usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

/*
Função chamada pela rota GET /agendamentos/propriedades/{propriedadeId}
- Rota de uso do ADM

O que faz:
  - Verifica se o usuário logado é admin, se não for, bloqueia o acesso
  - Recupera o ID da propriedade que esta na URL
  - Chama a função que busca todos os agendamentos de uma propriedade baseado no ID recuperado
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 200
  - lista de agendamentos da propriedade

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarAgendamentosPropriedade(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("você não tem permissão para realizar essa ação"))
		return
	}

	parametros := mux.Vars(r)
	propriedadeID, erro := strconv.ParseUint(parametros["propriedadeId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var agendamentos []models.Agendamento

	repo := repositorios.NovoRepoAgendamento(db)
	agendamentos, erro = repo.BuscarAgendamentosPropriedade(propriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

/*
Função chamada pela rota GET /agendamentos/{data}
- Rota de uso do ADM

O que faz:
  - Verifica se o usuário logado é admin, se não for, bloqueia o acesso
  - Recupera a data desejada que esta na URL
  - Data formato: AAAA-MM-DD
  - Formatos aceitos:
  - /agendamento/ano
  - /agendamento/ano-mes
  - /agendamento/ano-mes-dia
  - Verifica se o formato da data esta correto
  - Caso ok, chama a função que busca todos os agendamentos baseado no período especificado pela data
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 200
  - lista de agendamentos do período

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarAgendamentosPorData(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("você não tem permissão para realizar essa ação"))
		return
	}

	parametros := mux.Vars(r)
	dataParametro := parametros["data"]

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var data models.Data

	data, erro = data.VerificaData(dataParametro)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	var calendario []models.Calendario

	repo := repositorios.NovoRepoCalendario(db)
	calendario, erro = repo.BuscarAgendamentosPorData(data)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if calendario == nil {
		respostas.JSONresponse(w, http.StatusNoContent, nil)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, calendario)
}

/*
Função chamada pela rota GET /agendamentos/usuario/{data}
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera o ID do usuário através do token
  - Recupera a data desejada que esta na URL
  - Data formato: AAAA-MM-DD
  - Formatos aceitos:
  - /agendamento/ano
  - /agendamento/ano-mes
  - /agendamento/ano-mes-dia
  - Verifica se o formato da data esta correto
  - Caso ok, chama a função que busca todos os agendamentos do usuário logado baseado no período especificado pela data
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 200
  - lista de agendamentos do período

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarAgendamentosPorDataLogado(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("eu sei que você é admin e pode fazer tudo, mas essa rota é exclusiva do cliente"))
		return
	}

	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	dataParametro := parametros["data"]

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var data models.Data

	data, erro = data.VerificaData(dataParametro)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	var agendamentos []models.Calendario

	repo := repositorios.NovoRepoCalendario(db)
	agendamentos, erro = repo.BuscarAgendamentosPorDataLogado(data, usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

/*
Função chamada pela rota GET /agendamentos/adm/{usuarioId}
- Rota de uso do ADM

O que faz:
  - Verifica se o usuário logado é admin, se não for, bloqueia o acesso
  - Recupera a o ID do usuário desejado que esta na URL
  - Caso ok, chama a função que busca todos os agendamentos do usuário através do ID recuperado
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 200
  - lista de agendamentos do usuário

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarAgendamentosUsuarioId(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("você não tem permissão para realizar essa ação"))
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var agendamentos []models.Calendario

	repo := repositorios.NovoRepoCalendario(db)
	agendamentos, erro = repo.BuscarAgendamentosDoUsuario(usuarioId)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

/*
Função chamada pela rota POST /agendamentos/propriedades/{propriedadeId}
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Recupera o ID da propriedade desejada pela URL
  - Lê a request com as informações do novo agendamento
  - Formato da data do campo "diaAgendamento"
  - DD-MM-AAAA
  - D-M-AA
  - A sequência IMPORTA
  - Formato da hora do campo "checkout/checkin"
  - HH:MM:SS
  - "check": "15"
  - Apenas um valor ele entende 00:00:SS
  - "check": "15:30"
  - A partir de 2 valores ele ja entende HH:MM:00
  - Caso ok, chama a função que busca a propriedade do agendamento baseada no ID da URL
  - Verifica se a propriedade encontrada pertence ao usuário logado
  - Caso ok, chama a função que cria um agendamento com os dados fornecidos
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 201
  - ID do agendamento criado

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func AdicionarAgendamento(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	propriedadeID, erro := strconv.ParseUint(parametros["propriedadeId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var agendamento models.Agendamento
	if erro = json.Unmarshal(corpoRequest, &agendamento); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	agendamento.PropriedadeID = propriedadeID

	if erro = agendamento.Preparar("criando"); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repoPropriedade := repositorios.NovoRepoPropriedade(db)
	dbPropriedade, erro := repoPropriedade.BuscarPropriedadePorId(agendamento.PropriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		usuarioId, erro := auth.PegaUsuarioIDToken(r)
		if erro != nil {
			respostas.JSONerror(w, http.StatusUnauthorized, erro)
			return
		}
		if dbPropriedade.ProprietarioID != usuarioId {
			respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido criar agendamento para uma propriedade que não é a sua"))
			return
		}
	} 

	repoAgendamento := repositorios.NovoRepoAgendamento(db)
	agendamento.ID, erro = repoAgendamento.CriarAgendamento(agendamento)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusCreated, agendamento.ID)
}

/*
Função chamada pela rota PATCH /agendamentos/{agendamentoId}
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Recupera o ID do agendamento que ele quer atualizar pela URL
  - Chama a função que busca o cliente ID a partir do agendamento ID
  - Verifica se o usuário logado é o dono do agendamento
  - Lê a request com as informações da atualização
  - Formato da data do campo "diaAgendamento"
  - DD-MM-AAAA
  - D-M-AA
  - A sequência IMPORTA
  - Formato da hora do campo "checkout/checkin"
  - HH:MM:SS
  - "check": "15"
  - Apenas um valor ele entende 00:00:SS
  - "check": "15:30"
  - A partir de 2 valores ele ja entende HH:MM:00
  - Caso ok, chama a função que atualiza o agendamento baseado no ID da URL
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 204

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func AtualizarAgendamento(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	agendamentoID, erro := strconv.ParseUint(parametros["agendamentoId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoAgendamento(db)
	clienteID, erro := repo.BuscarClientePorAgendamentoId(agendamentoID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		usuarioID, erro := auth.PegaUsuarioIDToken(r)
		if erro != nil {
			respostas.JSONerror(w, http.StatusUnauthorized, erro)
			return
		}
		if clienteID != usuarioID {
			respostas.JSONerror(w, http.StatusForbidden, errors.New("você não tem permissão para atualizar um agendamento de outro usuário"))
			return
		}
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var agendamento models.Agendamento
	if erro = json.Unmarshal(corpoRequest, &agendamento); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = agendamento.Preparar("atualizando"); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = repo.AtualizarAgendamento(agendamentoID, agendamento); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

/*
Função chamada pela rota DELETE /agendamentos/{agendamentoId}
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Recupera o ID do agendamento que ele quer deletar pela URL
  - Chama a função que busca o cliente ID a partir do agendamento ID
  - Verifica se o usuário logado é o dono do agendamento que será deletado
  - Lê a request com as informações da atualização
  - Formato da data do campo "diaAgendamento"
  - DD-MM-AAAA
  - D-M-AA
  - A sequência IMPORTA
  - Formato da hora do campo "checkout/checkin"
  - HH:MM:SS
  - "check": "15"
  - Apenas um valor ele entende 00:00:SS
  - "check": "15:30"
  - A partir de 2 valores ele ja entende HH:MM:00
  - Caso ok, chama a função que deleta o agendamento especificado na URL
  - Retorna um caso de sucesso ou um caso de fracasso.

- Sucesso:
  - status code 204

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func RemoverAgendamento(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	agendamentoID, erro := strconv.ParseUint(parametros["agendamentoId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//verificar se o agendamento existe
	var agendamento models.Agendamento
	repo := repositorios.NovoRepoAgendamento(db)
	agendamento, erro = repo.BuscarAgendamentoPorId(agendamentoID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	
	if (agendamento.ID == 0) {
		respostas.JSONerror(w, http.StatusInternalServerError, errors.New("agendamento não encontrado"))
		return
	}

	repo = repositorios.NovoRepoAgendamento(db)
	clienteID, erro := repo.BuscarClientePorAgendamentoId(agendamentoID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		usuarioID, erro := auth.PegaUsuarioIDToken(r)
		if erro != nil {
			respostas.JSONerror(w, http.StatusUnauthorized, erro)
			return
		}
		if clienteID != usuarioID {
			respostas.JSONerror(w, http.StatusForbidden, errors.New("você não tem permissão de remover um agendamento de outro usuário"))
			return
		}
	}

	if erro = repo.DeletarAgendamento(agendamentoID); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}
