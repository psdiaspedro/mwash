package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONresponse(w http.ResponseWriter, statusCode int, dado interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dado != nil {
		if erro := json.NewEncoder(w).Encode(dado); erro != nil {
			log.Fatal(erro)
		}
	}
}

func JSONerror(w http.ResponseWriter, statusCode int, erro error) {
	JSONresponse(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}