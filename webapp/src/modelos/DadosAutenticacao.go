package modelos

//DadosAutenticacao contém o ID e o Token autenticados
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
