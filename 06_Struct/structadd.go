/* colecoes */
package main

import (
	"fmt"
)

type PessoaDados struct {
	cpf   string
	nome  string
	email string
}

type Pessoas struct {
	pessoa []PessoaDados
}

func main() {
	var count int
	var pessoa PessoaDados
	cpessoas := Pessoas{}
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

		cpessoas.Add(pessoa)
		count = count + 1

	}
	fmt.Printf("%v", cpessoas)
}

func (pessoas *Pessoas) Add(pessoa PessoaDados) {
	pessoas.pessoa = append(pessoas.pessoa, pessoa)
}
