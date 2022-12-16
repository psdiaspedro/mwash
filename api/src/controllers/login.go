package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuarioRequest models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuarioRequest); erro != nil {
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
	usuarioBanco, erro := repo.BuscarUsuarioPorEmail(usuarioRequest.Email)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}
	
	if erro = seguranca.ChecaSenha(usuarioRequest.Senha, usuarioBanco.Senha); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
	}

	usuarioBanco.Token, erro = auth.GerarToken(usuarioBanco.ID, usuarioBanco.Admin)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
	}

	respostas.JSONresponse(w, http.StatusOK, struct {
		ID			uint64	`json:"id,omitempty"`
		Nome		string	`json:"nome,omitempty"`
		Token		string	`json:"token,omitempty"`
	}{
		ID: usuarioBanco.ID,
		Nome: usuarioBanco.Nome,
		Token: usuarioBanco.Token,
	})
}