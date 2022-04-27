package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
)

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

func main() {

	/* carrega os dados da pessoa */
	pessoa := pessoaREAD()

	/* conecta no bd */
	isOk := mydb.MyDbConnect()
	if isOk {
		/* executa a querie */
		deletedRows, err := pessoaDELETE(pessoa)
		if err != nil {
			log.Fatal("Error deleting Employee: ", err.Error())
		}
		fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
	} else {
		fmt.Printf("Not Connected! %v", isOk)
	}
}

func pessoaREAD() Pessoa {
	var pessoa Pessoa
	fmt.Print("cpf a remover:")
	fmt.Scan(&pessoa.Cpf)
	return pessoa
}

func pessoaDELETE(pessoa Pessoa) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM Pessoa WHERE PessoaCpf = @PessoaCpf;")

	// Execute non-query with named parameters
	result, err := mydb.Db.ExecContext(ctx, tsql, sql.Named("PessoaCpf", &pessoa.Cpf))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
