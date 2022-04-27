package main

import (
	"fmt"
)

type Pessoas struct {
	Cpf   string
	Nome  string
	Email string
}

func main() {
	pessoas := []Pessoas{
		{"111", "John Doe", "gardener@email.com"},
		{"222", "Roger Roe", "driver@email.com"},
		{"333", "Paul Smith", "programmer@email.com"},
	}

	for _, pessoa := range pessoas {
		fmt.Printf("cpf:%v,nome:%v,email:%v\n", pessoa.Cpf, pessoa.Nome, pessoa.Email)
	}
}
