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

	fmt.Printf("coleção: %v", pessoas)
}
