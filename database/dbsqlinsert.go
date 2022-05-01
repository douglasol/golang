package main

import (
	"database/sql"
	"fmt"

	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
)

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

func main() {
	pessoa := pessoaGET()
	if mydb.MyDbConnect() {
		total, _ := pessoaINSERT(pessoa)
		if total > 0 {
			fmt.Printf("Inserção executada, %v", total)
		}
	} else {
		fmt.Printf("Não conectado!")
	}
}

func pessoaGET() Pessoa {
	var pessoa Pessoa
	fmt.Println("INSERT")
	fmt.Print("cpf:")
	fmt.Scan(&pessoa.Cpf)
	fmt.Print("nome:")
	fmt.Scan(&pessoa.Nome)
	fmt.Print("email:")
	fmt.Scan(&pessoa.Email)
	return pessoa
}

func pessoaINSERT(pessoa Pessoa) (int64, error) {
	err, ctx := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}

	tsql := `insert into pessoa (PessoaCPF,PessoaNome,PessoaEmail) values (@PessoaCPF,@PessoaNome,@PessoaEmail);`
	stmt, err := mydb.Db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	//row := stmt.QueryRowContext(
	row := stmt.QueryRow(
		ctx,
		sql.Named("PessoaCPF", pessoa.Cpf),
		sql.Named("PessoaNome", pessoa.Nome),
		sql.Named("PessoaEmail", pessoa.Email),
	)
	err = row.Scan()
	if err == nil {
		return -1, err
	}
	return 1, err
}
