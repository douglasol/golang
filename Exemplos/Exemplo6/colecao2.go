package main

import (
	"fmt"
)

type Pessoa struct {
	cpf   string
	nome  string
	email string
}

func main() {
	var count int
	var pessoa Pessoa
	pessoas := []Pessoa{}

	fmt.Println("Entre com dois registros na estrutura:")

	for {
		if count == 2 {
			break
		}

		fmt.Print("cpf:")
		fmt.Scan(&pessoa.cpf)

		fmt.Print("nome:")
		fmt.Scan(&pessoa.nome)

		fmt.Print("email:")
		fmt.Scan(&pessoa.email)

		pessoas = append(pessoas, pessoa)
		count = count + 1

	}
	fmt.Printf("%v", pessoas)
}
