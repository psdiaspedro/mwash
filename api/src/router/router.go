package router

import (
	"api/src/router/rotas"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI			string
	Method		string
	Function	func(w http.ResponseWriter, r *http.Request)
	NeedAuth	bool
}

func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigurarRotas(r)
}