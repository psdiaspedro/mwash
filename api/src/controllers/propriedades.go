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

/*
	Função chamada pela rota POST /minhas_propriedades

	O que faz:
		- Recupera o ID do usário logado através do token
		- Lê a request com as informações da nova propriedade
		- Se tudo estiver OK, chama a função que cria uma propriedade no banco de dados
		- Retorna um caso de sucesso ou um caso de fracasso.

	- Sucesso:
		- status code 201
		- ID da propriedade criada
	- Fracasso:
		- Retorna algum status code negativo
		- Retorna o erro de acordo com o problema
*/
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
	
	if erro = propriedade.Preparar("cadastrar"); erro != nil {
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

/*
	Função chamada pela rota GET /minhas_propriedades

	O que faz:
		- Recupera o ID do usário logado através do token
		- Se tudo estiver OK, chama a função que busca as propriedades do ID recuperado.
		- Retorna um caso de sucesso ou um caso de fracasso.

	- Sucesso:
		- status code 200
		- JSON com uma lista de todas as propriedades do ID logado
			- null, caso não encontre nada
	- Fracasso:
		- Retorna algum status code negativo
		- Retorna o erro de acordo com o problema
*/
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

/*
	Função chamada pela rota PATCH /minhas_propriedades/{propriedadeId}

	O que faz:
		- Recupera o ID do usário logado através do token
		- Recupera a propriedadeID do do parametro da URL
		- Busca no database uma propriedade baseada no ID
		- Verifica se a propriedade existe
		- Verifica se a propriedade é do usário logado
		- Lê a request com as informações da nova propriedade
		- Se tudo estiver OK, chama a função que atualiza a propriedade no banco de dados
		- Retorna um caso de sucesso ou um caso de fracasso.

	- Sucesso:
		- Status code 204
		- Lista com as propriedades do usuário logado
	- Fracasso:
		- Retorna algum status code negativo
		- Retorna o erro de acordo com o problema
*/
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

	if erro = dbPropriedade.PropriedadeCadastrada(); erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
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

	if erro = propriedade.Preparar("atualizar"); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = repo.AtualizarPropriedade(propriedadeID, propriedade); erro != nil {
		respostas.JSONerror(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSONresponse(w, http.StatusNoContent, nil)
}

/*
	Função chamada pela rota PATCH /minhas_propriedades/{propriedadeId}

	O que faz:
		- Recupera o ID do usário logado através do token
		- Recupera a propriedadeID do do parametro da URL
		- Busca no database uma propriedade baseada no ID
		- Verifica se a propriedade existe
		- Verifica se a propriedade é do usário logado
		- Lê a request com as informações da nova propriedade
		- Se tudo estiver OK, chama a função que deleta a propriedade no banco de dados
		- Retorna um caso de sucesso ou um caso de fracasso.

	- Sucesso:
		- status code 204
	- Fracasso:
		- Retorna algum status code negativo
		- Retorna o erro de acordo com o problema
*/
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

	if erro = dbPropriedade.PropriedadeCadastrada(); erro != nil {
		respostas.JSONerror(w, http.StatusUnprocessableEntity, erro)
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