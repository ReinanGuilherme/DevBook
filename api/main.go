package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//carregando as configurações de ambiente
	config.Carregar()

	r := router.Gerar()

	fmt.Println("Escutando na porta 5000")
	//passando valor da porta atraves do config
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortaServidor), r))
}
