package main

import (
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
	isOk := mydb.MyDbConnect()
	if isOk {
		var pessoas []Pessoa
		count, err, pessoas := pessoaSELECT()
		if err != nil {
			log.Fatal("Erro em Pessoa: ", err.Error())
		}
		fmt.Printf("Lidos %d registro(s) com sucesso.\n", count)
		for _, pessoa := range pessoas {
			fmt.Printf("%v\t%v\t%v\n", pessoa.Cpf, pessoa.Nome, pessoa.Email)
		}
	} else {
		fmt.Printf("Não conectado %v", isOk)
	}
}

func pessoaSELECT() (int, error, []Pessoa) {
	var count int
	var pessoas []Pessoa
	err, _ := mydb.MyDbContext()
	if err != nil {
		return -1, err, nil
	}
	tsql := "select PessoaCPF,PessoaNome,PessoaEmail from Pessoa;"
	rows, err := mydb.Db.Query(tsql)
	if err != nil {
		return -1, err, pessoas
	}
	defer rows.Close()
	for rows.Next() {
		var pessoa Pessoa
		err := rows.Scan(&pessoa.Cpf, &pessoa.Nome, &pessoa.Email)
		if err != nil {
			return -1, err, pessoas
		}
		pessoas = append(pessoas, pessoa)
		count++
	}
	return count, nil, pessoas
}
