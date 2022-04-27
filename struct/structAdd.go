package main

import (
	"fmt"
)

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

type Pessoas struct {
	pessoa []Pessoa
}

func main() {
	var count int
	var pessoa Pessoa
	pessoas := Pessoas{}
	fmt.Println("Entre com dois registros:")
	for {

		if count == 2 {
			break
		}

		fmt.Print("cpf:")
		fmt.Scan(&pessoa.Cpf)
		fmt.Print("nome:")
		fmt.Scan(&pessoa.Nome)
		fmt.Print("email:")
		fmt.Scan(&pessoa.Email)

		pessoas.Add(pessoa)
		count++

	}
	fmt.Printf("%v", pessoas)
}

func (pessoas *Pessoas) Add(pessoa Pessoa) {
	pessoas.pessoa = append(pessoas.pessoa, pessoa)
}
