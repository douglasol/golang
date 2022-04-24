package main

import (
	"fmt"
	"strings"
)

var Nomes []string
var NomesUpper []string
var NomesLower []string

func main() {
	nomesListGET2()

	for _, nome := range Nomes {
		fmt.Printf("%v\n", nome)
	}
	for _, nome := range NomesUpper {
		fmt.Printf("%v\n", nome)
	}
	for _, nome := range NomesLower {
		fmt.Printf("%v\n", nome)
	}
}

func nomesListGET2() {
	Nomes = append(Nomes, "Zé")
	Nomes = append(Nomes, "João")
	Nomes = append(Nomes, "Maria")

	// conversao maiusculo
	for _, nome := range Nomes {
		NomesUpper = append(NomesUpper, strings.ToUpper(nome))
	}
	// conversao minusculo
	for _, nome := range Nomes {
		NomesLower = append(NomesLower, strings.ToLower(nome))
	}
}
