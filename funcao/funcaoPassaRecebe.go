package main

import (
	"fmt"
	"strings"
)

func main() {
	nomes := nomesGET()
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}

	nomes = nomesToUpperGET(nomes)
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
	nomes = nomesToLowerGET(nomes)
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
}

func nomesGET() []string {
	nomes := []string{}
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	return nomes
}

func nomesToUpperGET(nomes []string) []string {
	nomesupper := []string{}
	for _, nome := range nomes {
		nomesupper = append(nomesupper, strings.ToUpper(nome))
	}
	return nomesupper
}
func nomesToLowerGET(nomes []string) []string {
	nomeslower := []string{}
	for _, nome := range nomes {
		nomeslower = append(nomeslower, strings.ToLower(nome))
	}
	return nomeslower
}
