package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*
Para carregar as variáveis ambiente vamos usar um pacote externo: go get github.com/joho/godotenv
*/
var (
	//APIURL representa a URL para comunicação com a API
	APIURL = ""

	//Porta número da porta onde a API vai estar rodando
	Porta = 0

	//HashKey é utilizada para autenticar o cookie
	HashKey []byte

	//BlockKey é utilizada para cripitografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT")) //strconv.Atoi converte uma string em númerico
	if erro != nil {
		Porta = 9300
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
