package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI        string
	Metodo     string
	Funcao     func(w http.ResponseWriter, r *http.Request)
	RequerAuth bool
}

func ConfigurarRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasPropriedades...)
	rotas = append(rotas, rotasAgendamentos...)
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotaCadastro)
	rotas = append(rotas, rotasAgendamentos...)
	rotas = append(rotas, rotaAuth)

	for _, rota := range rotas {

		if rota.RequerAuth {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Auth(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}
	}
	return r
}
