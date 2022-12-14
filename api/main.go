package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarEnv()
	r := router.GerarRotas()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}