package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	/*
		//Este trecho de código serve para criar uma chave secreta aleatória
		//com base 64 para ser usada com SECRET na geração do TOKEN
		chave := make([]byte, 64)
		if _, erro := rand.Read(chave); erro != nil {
			log.Fatal(erro)
		}
		stringB64 := base64.StdEncoding.EncodeToString(chave)
		fmt.Println(stringB64)
	*/

}

func main() {
	config.Carregar()

	fmt.Println("PORTA:", config.Porta)
	fmt.Println("CONEXAO:", config.StringConexaoBanco)
	fmt.Println("Rodando API!")

	r := router.Gerar()

	fmt.Println("Escutando na porta :", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
