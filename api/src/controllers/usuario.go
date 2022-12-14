package controllers

import (
	"net/http"
)

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario..."))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario..."))
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Senha..."))
}