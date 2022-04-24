package main

import "fmt"

func tConcatString() {
	var nome string
	var sobrenome string

	nome = "Zé"
	sobrenome = "Silva"
	var nomeinteiro = nome + " " + sobrenome
	fmt.Println(nomeinteiro)
}
