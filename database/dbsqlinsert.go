package main

import (
	"context"
	"database/sql"
	"errors"
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
		count, err := pessoaINSERT(pessoa)
		if err != nil {
			log.Fatal("Error reading Pessoas: ", err.Error())
		}
		fmt.Printf("Ok.%v\n", count)

	} else {
		fmt.Printf("Not Connected! %v", isOk)
	}
}

func pessoaREAD() Pessoa {
	var pessoa Pessoa
	fmt.Print("cpf:")
	fmt.Scan(&pessoa.Cpf)
	fmt.Print("nome:")
	fmt.Scan(&pessoa.Nome)
	fmt.Print("email:")
	fmt.Scan(&pessoa.Email)
	return pessoa
}

func pessoaINSERT(pessoa Pessoa) (int64, error) {
	ctx := context.Background()
	var err error

	if mydb.Db == nil {
		err = errors.New("CreateEmployee: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
		insert into pessoa (PessoaCPF,PessoaNome,PessoaEmail) values (@PessoaCPF,@PessoaNome,@PessoaEmail);
		select isNull(SCOPE_IDENTITY(), -1);
		`

	stmt, err := mydb.Db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("PessoaCPF", pessoa.Cpf),
		sql.Named("PessoaNome", pessoa.Nome),
		sql.Named("PessoaEmail", pessoa.Email))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}
