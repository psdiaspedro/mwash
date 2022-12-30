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

func AdicionarPropriedade(w http.ResponseWriter, r *http.Request) {
	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var propriedade models.Propriedade
	if erro = json.Unmarshal(corpoRequest, &propriedade); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	propriedade.ProprietarioID = usuarioIdToken
	
	if erro = propriedade.Preparar(); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoPropriedade(db)
	propriedade.ID, erro = repo.CriarPropriedade(propriedade)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusCreated, propriedade.ID)
}

func ListarPropriedades(w http.ResponseWriter, r *http.Request) {
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

	var propriedades []models.Propriedade
	
	repo := repositorios.NovoRepoPropriedade(db)
	propriedades, erro = repo.BuscarPropriedadesDoUsuario(usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, propriedades)
}

func AtualizarPropriedade(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.PegaUsuarioIDToken(r)
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

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoPropriedade(db)
	dbPropriedade, erro := repo.BuscarPropriedadePorId(propriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if dbPropriedade.ProprietarioID != usuarioID {
		respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido atualizar propriedade de outro usuário"))
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var propriedade models.Propriedade
	if erro = json.Unmarshal(corpoRequest, &propriedade); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = propriedade.Preparar(); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = repo.AtualizarPropriedade(propriedadeID, propriedade); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

func RemoverPropriedade(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.PegaUsuarioIDToken(r)
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

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoPropriedade(db)
	dbPropriedade, erro := repo.BuscarPropriedadePorId(propriedadeID)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if dbPropriedade.ProprietarioID != usuarioID {
		respostas.JSONerror(w, http.StatusForbidden, errors.New("proibido deletar propriedade de outro usuário"))
		return
	}

	if erro = repo.DeletarPropriede(propriedadeID); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}