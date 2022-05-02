package main

import (
	"context"
	"encoding/json"
	"exemplo9/mydb"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	/* conecta no bd */
	isOk := mydb.MyDbConnect()
	if isOk {
		count, err, pessoas := pessoaGET()
		if err != nil {
			log.Fatal("Erro lendo Pessoa: ", err.Error())
		}

		/* converte struct em JSON */
		data, err := json.Marshal(pessoas)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			os.Stdout.Write(data)
		}
		fmt.Printf("%d registros recuperadas.\n", count)

	} else {
		fmt.Printf("Não conectado! %v", isOk)
	}
}

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

func pessoaGET() (int, error, []Pessoa) {

	var count int
	var pessoa Pessoa
	pessoas := []Pessoa{}

	/* Checa se o banco está vivo */
	ctx := context.Background()
	err := mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err, nil
	}

	/* Execute query no banco */
	tsql := "select top 10 pessoaCPF, pessoaNome, pessoaEmail from pessoa;"
	rows, err := mydb.Db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err, nil
	}
	defer rows.Close()

	/* percorre os registros recuperados criando uma coleção */
	for rows.Next() {
		var pessoacpf, pessoaNome, pessoaEmail string
		err := rows.Scan(&pessoacpf, &pessoaNome, &pessoaEmail)
		if err != nil {
			return -1, err, nil
		}
		/* com remoção de espaços */
		pessoa.Cpf = strings.TrimSpace(pessoacpf)
		pessoa.Nome = strings.TrimSpace(pessoaNome)
		pessoa.Email = strings.TrimSpace(pessoaEmail)
		pessoas = append(pessoas, pessoa)
		count++
	}

	return count, nil, pessoas
}
