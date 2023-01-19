package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

/*
	Função chamada pela rota /login

	O que faz:
		- Lê a request de login, faz análises, validações e verificações de segurança
		- Se tudo estiver OK, chama a função que busca um usuário no banco de dados
		- Rtorna um caso de sucesso ou um caso de fracasso.

	- Sucesso:
		- status code 200
		- JSON com ID, nome do usuário e token string JWT
			- Token contém:
				- Exp - 12h
				- ID do usuário
				- Informação se é admin ou não (boolean)
	- Fracasso:
		- Retorna algum status code negativo
		- Retorna o erro de acordo com o problema
*/
func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuarioRequest models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuarioRequest); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro =  usuarioRequest.Preparar("login"); erro != nil {
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
	
	if erro =  usuarioBanco.UsuarioCadastrado(); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	if erro = seguranca.ChecaSenha(usuarioRequest.Senha, usuarioBanco.Senha); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
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