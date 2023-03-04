package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/*
Função chamada pela rota GET /usuario
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Caso ok, chama a função que busca as informações do usuário logado
  - Retorna um caso de sucesso ou um caso de fracasso

- Sucesso:
  - status code 200
  - Informações do usuário

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func BuscarDadosUsuario(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("eu sei que você é admin e pode fazer tudo, mas essa rota é exclusiva do cliente"))
		return
	}

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

	repo := repositorios.NovoRepoUsuario(db)
	usuario, erro := repo.BuscarUsuarioPorId(usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, usuario)
}

func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	
	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if !isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("Rota exclusiva do admin"))
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuario(db)
	clientes, erro := repo.BuscarClientes()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, clientes)
}

/*
Função chamada pela rota PATCH /usuario/atualizar_dados
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Lê a request com os dados de atualização
  - Faz validações com os dados lidos
  - Caso ok, chama a função que atualiza as informações do usuário logado
  - Retorna um caso de sucesso ou um caso de fracasso

- Sucesso:
  - status code 204

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("eu sei que você é admin e pode fazer tudo, mas essa rota é exclusiva do cliente"))
		return
	}

	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

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

	if erro = usuario.Preparar("atualizar"); erro != nil {
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
	if erro = repo.AtualizarDadosUsuario(usuarioIdToken, usuario); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

/*
Função chamada pela rota PATCH /usuario/senha
- Rota de uso do cliente

O que faz:
  - Verifica se o usuário logado é admin, se for, bloqueia o acesso
  - Recupera a o ID do usuário logado pelo token
  - Lê a request com os dados de atualização da senha
  - Faz validações com os dados lidos
  - Caso ok, chama a função que busca a senha atual do usuário logado no database
  - Verifica se a senha informada bate com a do database
  - Transforma a senha string em um hash
  - Caso ok, chama a função que atualiza a senha do usuário logado
  - Retorna um caso de sucesso ou um caso de fracasso

- Sucesso:
  - status code 204

- Fracasso:
  - Retorna algum status code negativo
  - Retorna o erro de acordo com o problema
*/
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	isAdmin, erro := auth.IsAdmin(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	} else if isAdmin {
		respostas.JSONerror(w, http.StatusUnauthorized, errors.New("eu sei que você é admin e pode fazer tudo, mas essa rota é exclusiva do cliente"))
		return
	}

	usuarioIdToken, erro := auth.PegaUsuarioIDToken(r)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var senha models.Senha
	if erro = json.Unmarshal(corpoRequest, &senha); erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = senha.Preparar(); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.ConectarBancoDeDados()
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuario(db)
	senhaAtual, erro := repo.BuscarSenhaAtualUsuario(usuarioIdToken)
	if erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.ChecaSenha(senha.Atual, senhaAtual); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	senhaComHash, erro := seguranca.GerarHash(senha.Nova)
	if erro != nil {
		respostas.JSONerror(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repo.AtualizarSenha(usuarioIdToken, string(senhaComHash)); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

func ValidarToken(w http.ResponseWriter, r *http.Request) {
	if erro := auth.ValidaToken(r); erro != nil {
		respostas.JSONerror(w, http.StatusUnauthorized, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusOK, nil)
}
