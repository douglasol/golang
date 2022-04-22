package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	email     string
	idade     uint
}

func main() {
	var pessoas []Pessoa

	pessoas = append(pessoas, pessoaADD("Zé", "Mané", "ze@ze.com", 10))
	pessoas = append(pessoas, pessoaADD("Maria", "José", "zeze@zeze.com", 10))
	for _, pessoa := range pessoas {
		fmt.Printf("%v, %v, %v, %v\n", pessoa.nome, pessoa.sobrenome, pessoa.email, pessoa.idade)
	}
}

func pessoaADD(nome string, sobrenome string, email string, idade uint) Pessoa {
	var pessoa = Pessoa{
		nome:      nome,
		sobrenome: sobrenome,
		email:     email,
		idade:     idade,
	}
	return pessoa
}
