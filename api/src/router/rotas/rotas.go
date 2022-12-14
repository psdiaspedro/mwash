package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI			string
	Metodo		string
	Funcao		func(w http.ResponseWriter, r *http.Request)
	RequerAuth	bool
}

func ConfigurarRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasPropriedades...)
	rotas = append(rotas, rotasAgendamentos...)
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotaCadastro)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}