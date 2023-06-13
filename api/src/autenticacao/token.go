package autenticacao

/*
	JSON web token gera um token de todas as informações enviada via web
	é necessário instalação do pacote externo: go get github.com/dgrijalva/jwt-go
*/

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go" // o primeiro parâmetro é um alias pois o pacote contém traço -
)

func CriarToken(usuarioID int64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["autorized"] = true                                 //autorizado/não autorizado
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()       //tempo para expirar (6h) esse .Unix() é o $horolog do MUMPS
	permissoes["UsuarioID"] = usuarioID                            //ID do usuário
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) //para esse método de assinatura HS256 é recomendado usar uma SECRET de 64 bytes
	return token.SignedString([]byte(config.SecretKey))            //a geração do tokem apartir de uma SECRET_KEY criada no arquivo .env
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token Inválido!")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

// ExtrairUsuarioID retorna o UsuarioID que está no TOKEN
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["UsuarioID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("Token inválido!")
}
