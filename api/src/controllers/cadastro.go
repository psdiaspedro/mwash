package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

/*
	Função chamada pela rota /cadastrar

	O que faz:
		- Lê a request de cadastro, faz análises, validações e verificações de segurança
		- Se tudo estiver OK, chama a função que cadastra um usuário no banco de dados.
		- Retorna um caso de sucesso ou um caso de fracasso.

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

	corpoRequest, erro := io.ReadAll(r.Body)
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