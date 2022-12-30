package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarAgendamentos(w http.ResponseWriter, r *http.Request) {
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

	var agendamentos []models.Agendamento

	repo := repositorios.NovoRepoAgendamento(db)
	agendamentos, erro = repo.BuscarAgendamentosDoUsuario(usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

func BuscarAgendamentosPropriedade(w http.ResponseWriter, r *http.Request) {
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

func BuscarAgendamentosPorData(w http.ResponseWriter, r *http.Request) {
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

	var agendamentos []models.Agendamento

	repo := repositorios.NovoRepoAgendamento(db)
	agendamentos, erro = repo.BuscarAgendamentosPorData(data)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

func BuscarAgendamentosPorDataLogado(w http.ResponseWriter, r *http.Request) {
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

	var agendamentos []models.Agendamento

	repo := repositorios.NovoRepoAgendamento(db)
	agendamentos, erro = repo.BuscarAgendamentosPorDataLogado(data, usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

func BuscarAgendamentosUsuarioId(w http.ResponseWriter, r *http.Request) {
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

	var agendamentos []models.Agendamento

	repo := repositorios.NovoRepoAgendamento(db)
	agendamentos, erro = repo.BuscarAgendamentosDoUsuario(usuarioId)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSONresponse(w, http.StatusOK, agendamentos)
}

//cliente adicionando agendamento
func AdicionarAgendamento(w http.ResponseWriter, r *http.Request) {
	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}
	
	parametros := mux.Vars(r)
	propriedadeID, erro := strconv.ParseUint(parametros["propriedadeId"], 10, 64)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	
	corpoRequest, erro := ioutil.ReadAll(r.Body)
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

	if erro = agendamento.Preparar(); erro != nil {
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

	if dbPropriedade.ProprietarioID != usuarioIdToken {
		respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido criar agendamento para uma propriedade que não é a sua"))
		return
	}

	repoAgendamento := repositorios.NovoRepoAgendamento(db)
	agendamento.ID, erro = repoAgendamento.CriarAgendamento(agendamento) 
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusCreated, agendamento.ID)
}

//cliente atualizar agendamento
func AtualizarAgendamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

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

	//vai buscar o agendamento baseado no agendamentoID da URL
	repo := repositorios.NovoRepoAgendamento(db)
	dbAgendamento, erro := repo.BuscarAgendamentoPorId(agendamentoID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	//vai buscar a propriedade relacionada ao agendamento
	repoPropriedade := repositorios.NovoRepoPropriedade(db)
	dbPropriedade, erro := repoPropriedade.BuscarPropriedadePorId(dbAgendamento.PropriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	//vai verificar se o usuarioID é o mesmo do dono da propriedade
	if dbPropriedade.ProprietarioID != usuarioID {
		respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido atualizar um agendamento de outro usuário"))
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var agendamento models.Agendamento
	if erro = json.Unmarshal(corpoRequest, &agendamento); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = agendamento.Preparar(); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = repo.AtualizarAgendamento(agendamentoID, agendamento); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

//cliente remover agendamento
func RemoverAgendamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

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
	dbAgendamento, erro := repo.BuscarAgendamentoPorId(agendamentoID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	//vai buscar a propriedade relacionada ao agendamento
	repoPropriedade := repositorios.NovoRepoPropriedade(db)
	dbPropriedade, erro := repoPropriedade.BuscarPropriedadePorId(dbAgendamento.PropriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	//vai verificar se o usuarioID é o mesmo do dono da propriedade
	if dbPropriedade.ProprietarioID != usuarioID {
		respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido remover uma agendamento de outro usuário"))
		return
	}

	if erro = repo.DeletarAgendamento(agendamentoID); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}