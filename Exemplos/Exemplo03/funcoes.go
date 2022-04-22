package main

import "strings"

func alomamae() {
	println("alo mamae!")
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
