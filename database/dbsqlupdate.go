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
		// Update from database
		updatedRows, err := pessoaUPDATE(pessoa)
		if err != nil {
			log.Fatal("Error updating Employee: ", err.Error())
		}
		fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	} else {
		fmt.Printf("Not Connected! %v", isOk)
	}
}

func pessoaREAD() Pessoa {
	var pessoa Pessoa
	fmt.Print("cpf da pessoa:")
	fmt.Scan(&pessoa.Cpf)
	fmt.Print("nome novo:")
	fmt.Scan(&pessoa.Nome)
	fmt.Print("email novo:")
	fmt.Scan(&pessoa.Email)
	return pessoa
}

// UpdateEmployee updates an employee's information
func pessoaUPDATE(pessoa Pessoa) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE Pessoa SET PessoaNome=@PessoaNome, PessoaEmail=@PessoaEmail WHERE PessoaCpf = @PessoaCpf")

	// Execute non-query with named parameters
	result, err := mydb.Db.ExecContext(
		ctx,
		tsql,
		sql.Named("PessoaCpf", &pessoa.Cpf),
		sql.Named("PessoaNome", &pessoa.Nome),
		sql.Named("PessoaEmail", &pessoa.Email))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
