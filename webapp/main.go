package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	/*
		//Este trecho de código serve para criar uma chave secreta aleatória
		//com base 64 de 16 bytes para ser usada com SECRET na geração do TOKEN
		//de criação de cookies.
		//foi usado o pacote externo SECURECOOKIE
		hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
		blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
		fmt.Println(hashKey)
		fmt.Println(blockKey)
	*/
}

func main() {
	config.Carregar()
	cookies.Configurar()

	utils.CarregarTemplates()

	r := router.Gerar()
	fmt.Println("Escutando na porta: ", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
