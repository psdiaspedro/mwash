package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	
	// Le todo o corpo da request
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Cria uma variável baseada no modelo Usuário que contém todos os campos necessários
	// Parseia as infos da request para a variável
	var usuarioRequest models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuarioRequest); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	// Verifica se existe o campo email e senha, verifica se o email esta
	// no formato certo
	if erro =  usuarioRequest.Preparar("login"); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	// Abre conexão com o database
	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		fmt.Println(erro)
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Abre conexão com o repositório,
	// BuscarUsuarioPorEmail busca um possível usuário baseado no email fornecido
	repo := repositorios.NovoRepoUsuario(db)
	usuarioBanco, erro := repo.BuscarUsuarioPorEmail(usuarioRequest.Email)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}
	
	// Verifica se o retorno do banco foi um usuário válido
	if erro =  usuarioBanco.UsuarioCadastrado(); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	// Verifica se a senha fornecida bate com a do usuário que o banco retornou
	if erro = seguranca.ChecaSenha(usuarioRequest.Senha, usuarioBanco.Senha); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	// Gera o token com as infos
	// exp - expiração de 12 horas
	// usuarioID - ID do usuário logado
	// admin  - booleano com o tipo do usuário logado
	usuarioBanco.Token, erro = auth.GerarToken(usuarioBanco.ID, usuarioBanco.Admin)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
	}

	// Retorna status code 200,
	// Retorna um JSON com ID, Nome e Token do usuário logado
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