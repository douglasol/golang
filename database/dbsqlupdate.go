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
		total, err := pessoaUPDATE(pessoa)
		if err == nil {
			fmt.Printf("Atualizada %d registro com sucesso\n", total)
		}
	} else {
		fmt.Printf("Não conectado!")
	}
}

func pessoaGET() Pessoa {
	var pessoa Pessoa
	fmt.Println("UPDATE")
	fmt.Print("cpf:")
	fmt.Scanf("%s\n", &pessoa.Cpf)
	fmt.Print("nome:")
	fmt.Scanf("%s\n", &pessoa.Nome)
	fmt.Print("email:")
	fmt.Scanf("%s\n", &pessoa.Email)
	return pessoa
}

func pessoaUPDATE(pessoa Pessoa) (int64, error) {
	err, ctx := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}
	tsql := fmt.Sprintf("UPDATE Pessoa SET PessoaNome=@PessoaNome, PessoaEmail=@PessoaEmail WHERE PessoaCpf = @PessoaCpf")
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
