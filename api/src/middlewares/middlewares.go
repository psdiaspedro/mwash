package middlewares

import (
	"api/src/auth"
	"api/src/respostas"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidaToken(r); erro != nil {
			respostas.JSONerror(w, http.StatusUnauthorized, erro)
			return
		}
		
		next(w, r)
	}
}