package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar método para criar um novo usuário
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	fmt.Println("ultimo ID...")
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		fmt.Println("LAST ID:", erro)
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick do BD
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, CriadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			fmt.Println("erro no linhas.SCAN")
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscaPorID traz um usuário de uma ID no BD
func (repositorio Usuarios) BuscaPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM  usuarios WHERE id= ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	fmt.Println(usuario)

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera informações de um usuário no BD
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome=?, nick=?, email=? WHERE id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}

// Deleta um usuário no BD
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

// BuscarPorEmail busca um usuário por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE  email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Seguir cadastra um seguidor em um usuário do BD
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	//IGNORE serve para ignorar se já está seguindo, para não dar erro no INSERT
	statement, erro := repositorio.db.Prepare("INSERT IGNORE INTO seguidores (usuario_ID, seguidor_ID) VALUES (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// PararDeSeguir retira o cadastro de um seguidor em um usuário do BD
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	//IGNORE serve para ignorar se já está seguindo, para não dar erro no INSERT
	statement, erro := repositorio.db.Prepare("DELETE FROM seguidores  WHERE usuario_id=? AND seguidor_id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// BuscarSeguidores retorna seguidores de um usuário do BD
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`SELECT u.id, u.nome, u.nick, u.email, u.criadoEm 
	FROM usuarios u 
	INNER JOIN seguidores s on u.id=s.seguidor_id where s.usuario_id=?`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// AtualizarSenha retira o cadastro de um seguidor em um usuário do BD
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET senha=? WHERE id=? ")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}

// BuscarSenha retorna a senha de um usuário pelo ID do BD
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id=?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}
