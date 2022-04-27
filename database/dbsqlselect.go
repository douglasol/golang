/*  MS Sql server
 	Onde obter os drivers e instalação?
	https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go

*/
package main

import (
	"context"
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

	/* conecta no bd */
	isOk := mydb.MyDbConnect()
	if isOk {
		/* executa a querie */
		var pessoas []Pessoa
		count, err, pessoas := pessoaGET()
		if err != nil {
			log.Fatal("Error reading Pessoas: ", err.Error())
		}
		fmt.Printf("Read %d row(s) successfully.\n", count)
		for _, pessoa := range pessoas {
			fmt.Printf("%v\t%v\t%v\n", pessoa.Cpf, pessoa.Nome, pessoa.Email)
		}
	} else {
		fmt.Printf("Not Connected! %v", isOk)
	}
}

func pessoaGET() (int, error, []Pessoa) {
	var count int
	var pessoas []Pessoa

	/* Check if database is alive. */
	ctx := context.Background()
	err := mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err, pessoas
	}

	/* Execute query */
	tsql := "select PessoaCPF,PessoaNome,PessoaEmail from Pessoa;"
	rows, err := mydb.Db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err, pessoas
	}
	defer rows.Close()

	/* Iterate through the result set */
	for rows.Next() {
		var pessoa Pessoa
		/* Get values from row */
		err := rows.Scan(&pessoa.Cpf, &pessoa.Nome, &pessoa.Email)
		if err != nil {
			return -1, err, pessoas
		}
		pessoas = append(pessoas, pessoa)
		count++
	}
	return count, nil, pessoas
}
