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

/*
	Função chamada pela rota /cadastrar

	O que faz: 
		- Lê a request de cadastro, faz análises, validações e verificações de segurança
		- Se tudo estiver OK, chama a função que cadastra um usuário no banco de dados.
		- Retorna um caso de sucesso e um de fracasso.

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
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	
	// Le todo o corpo da request
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Criar uma verificaçäo do token
	// Só liberar o cadastro pra quem for ADMIN TRUE
	
	// Cria uma variável baseada no modelo Usuário que contém todos os campos necessários
	// Parseia as infos da request para a variável
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	// Trima os espaços das strings para evitar erros
	// Verifica se existe os campos obrigatórios (nome, email, senha, contato)
	// Verifica se o email esta no formato certo
	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	// Abre conexão com o database
	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Abre conexão com o repositório,
	// CadastrarUsuarioNoBanco cadastra o usuário no banco de dados
	repo := repositorios.NovoRepoUsuario(db)
	usuario.ID, erro = repo.CadastrarUsuarioNoBanco(usuario)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	// Retorna status code 201,
	// Retorna o ID do usuário criado
	respostas.JSONresponse(w, http.StatusCreated, usuario.ID)
}