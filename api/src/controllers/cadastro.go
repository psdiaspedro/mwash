package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}
	
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuario(db)
	usuario.ID, erro = repo.CadastrarUsuarioNoBanco(usuario)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusCreated, usuario.ID)
}