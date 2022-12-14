package controllers

import "net/http"

func BuscarAgendamento(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Agendamento..."))
}

func AdicionarAgendamento(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adicioando Agendamento..."))
}

func AtualizarAgendamento(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Atualizando Agendamento..."))
}

func RemoverAgendamento(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Removendo Agendamento..."))
}
