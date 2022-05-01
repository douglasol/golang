package main

import (
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
	pessoa := pessoaGET()
	isOk := mydb.MyDbConnect()
	if isOk {
		total, err := pessoaDELETE(pessoa)
		if err != nil {
			log.Fatal("Error deleting: ", err.Error())
		}
		fmt.Printf("Apagados %d registros com sucesso.\n", total)
	} else {
		fmt.Printf("Não conectado! %v", isOk)
	}
}

func pessoaGET() Pessoa {
	var pessoa Pessoa
	fmt.Println("DELETE")
	fmt.Print("cpf:")
	fmt.Scan(&pessoa.Cpf)
	return pessoa
}

func pessoaDELETE(pessoa Pessoa) (int64, error) {
	err, ctx := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}
	defer mydb.Db.Close()
	tsql := fmt.Sprintf("DELETE FROM Pessoa WHERE PessoaCpf = @PessoaCpf;")
	result, err := mydb.Db.ExecContext(ctx, tsql, sql.Named("PessoaCpf", &pessoa.Cpf))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
