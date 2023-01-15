package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	config.CarregarEnv()
	r := router.GerarRotas()
	h := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	m := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT"})
	o := handlers.AllowedOrigins([]string{"*"})

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handlers.CORS(h, m, o)(r)))
}
