package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty` //omitempty serve para quando os dados desse campo estiver vazio omitir o envio para JSON
	Nome     string    `json:nome,omitempty`
	Nick     string    `json:nick,omitempty`
	Email    string    `json:email,omitempty`
	Senha    string    `json:nome,omitempty`
	CriadoEm time.Time `json:CriadoEm,omitempty`
}

// Preparar chama os métodos para validar e validar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar(etapa)
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("Nome é obrigatorio e não pode estar me branco!")
	}
	if usuario.Nick == "" {
		return errors.New("Nick é obrigatorio e não pode estar me branco!")
	}
	if usuario.Email == "" {
		return errors.New("Email é obrigatório e não pode estar me branco!")
	}

	/*
		pacote externo que checa se o email é válido, necessário instalar : go get github.com/badoux/checkmail
	*/
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é invalido!")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha é obrigatorio e não pode estar me branco!")
	}

	return nil

}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
