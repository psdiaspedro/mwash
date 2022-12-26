package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
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

func AdicionarAgendamento(w http.ResponseWriter, r *http.Request) {
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

	repo := repositorios.NovoRepoAgendamento(db)
	agendamento.ID, erro = repo.CriarAgendamento(agendamento) 
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusCreated, agendamento.ID)
	w.Write([]byte("Adicioando Agendamento..."))
}

func AtualizarAgendamento(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Atualizando Agendamento..."))
}

func RemoverAgendamento(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Removendo Agendamento..."))
}
