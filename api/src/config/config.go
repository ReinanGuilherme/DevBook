package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexao é a string de conexão com o MySQL
	StringConexaoBanco = ""
	//Porta onde a API vai estar rodando
	PortaServidor = 0
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	//buscando porta do servidor
	PortaServidor, erro = strconv.Atoi(os.Getenv("portServer"))
	if erro != nil {
		PortaServidor = 9000
	}

	// variaveis de conexão
	var (
		host     = os.Getenv("host")
		port     = os.Getenv("port")
		database = os.Getenv("database")
		user     = os.Getenv("user")
		password = os.Getenv("password")
	)

	//string de conexão
	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", user, password, host, port, database)
}
