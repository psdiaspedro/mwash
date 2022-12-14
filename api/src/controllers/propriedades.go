package controllers

import "net/http"

func ListarPropriedades(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando Propriedades..."))
}

func AdicionarPropriedade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adicionando Propriedade..."))
}

func AtualizarPropriedade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Propriedade..."))
}

func RemoverPropriedade(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removendo Propriedade..."))
}