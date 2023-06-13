package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

/*
	necessário baixar o pacote que cria hash em uma string: go get golang.org/x/crypto/bcrypt
*/

//Hash recebe uma string e retorna seu HASH
func Hash(senha string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(senha),bcrypt.DefaultCost)

}

func VerificarSenha(senhaComHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}