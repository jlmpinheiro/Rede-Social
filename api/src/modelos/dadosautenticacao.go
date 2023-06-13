package modelos

//DadosAutenticacao contém o ID e o Token do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
