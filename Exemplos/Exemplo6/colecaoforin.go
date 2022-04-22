package main

import (
	"fmt"
)

type Pessoas struct {
	cpf   string
	nome  string
	email string
}

func main() {
	pessoas := []Pessoas{
		{"111", "John Doe", "gardener@email.com"},
		{"222", "Roger Roe", "driver@email.com"},
		{"333", "Paul Smith", "programmer@email.com"},
	}

	fmt.Printf("[\n")
	for _, pessoa := range pessoas {
		fmt.Printf("{\"cpf\":\"%v\",\t\"nome\":\"%v\",\t\"email\":\"%v\"},\n", pessoa.cpf, pessoa.nome, pessoa.email)
	}
	fmt.Printf("]\n")
}
