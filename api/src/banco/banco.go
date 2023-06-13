package banco

/*
	Tem que instalar neste diretório o driver do sql nesse diretorio: go get github.com/go-sql-driver/mysql
*/
import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o Banco de Dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		fmt.Println("OPEN:", erro)
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		fmt.Println("PING:", erro)
		db.Close()
		return nil, erro
	}

	return db, nil

}
