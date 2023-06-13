package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID         uint64       `json: "id"`
	Nome       string       `json: "nome"`
	Email      string       `json: "email"`
	Nick       string       `json: "nick"`
	CriadoEm   time.Time    `json: "criadoEM"`
	Seguidores []Usuario    `json:Seguidores`
	Seguindo   []Usuario    `json:segindo`
	Publicacao []Publicacao `json:publicacoes`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar os dados do usuário, essas requisições serão com canais de concorrência...
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				fmt.Println("Erro ao buscar o usuário")
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}
			//fmt.Println("usuarioCarregado:", usuarioCarregado)
			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				fmt.Println("Erro ao buscar seguidores")
				return Usuario{}, errors.New("Erro ao buscar seguidores")
			}
			//fmt.Println("seguidoresCarregados:", seguidoresCarregados)
			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				fmt.Println("Erro ao buscar quem está seguindo")
				return Usuario{}, errors.New("Erro ao buscar quem está seguindo")
			}
			//fmt.Println("seguindoCarregados:", seguindoCarregados)
			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				fmt.Println("Erro ao buscar publicações")
				return Usuario{}, errors.New("Erro ao buscar publicações")
			}
			//fmt.Println("publicacoesCarregadas:", publicacoesCarregadas)
			publicacoes = publicacoesCarregadas

		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacao = publicacoes

	return usuario, nil
}

// BuscarDadosDoUsuario chama a API para buscar os dados de cadastro do usuário
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil { //retorna um usuário nulo caso tenha algum problema
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuarios Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		canal <- Usuario{}
		return
	}
	//fmt.Println("usuarios:", usuarios)
	canal <- usuarios
}

// BuscarSeguidores chama a API para buscar quem o usuário está seguindo
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil { //retorna um usuário nulo caso tenha algum problema
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}
	//fmt.Println("seguidores:", seguidores)
	canal <- seguidores
}

// BuscarSeguindo chama a API para buscar quem está seguindo o usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	//fmt.Println("seguindo:", seguindo)
	canal <- seguindo
}

// BuscarPublicacoes chama a API para buscar as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}
	//.Println("publicacoes:", publicacoes)
	canal <- publicacoes
}
