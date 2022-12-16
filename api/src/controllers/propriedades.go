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
	w.Write([]byte("Adicionando Propriedade..."))
}

func ListarPropriedades(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando Propriedades..."))
}

func AtualizarPropriedade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Propriedade..."))
}

func RemoverPropriedade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removendo Propriedade..."))
}