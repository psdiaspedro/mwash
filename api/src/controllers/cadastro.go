package controllers

import "net/http"

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cadastrando Usuario..."))
}