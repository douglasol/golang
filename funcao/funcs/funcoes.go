package pck1

import "strings"

func NomesGET() []string {
	nomes := []string{}
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	return nomes
}

func NomesToUpperGET(nomes []string) []string {
	nomesupper := []string{}
	for _, nome := range nomes {
		nomesupper = append(nomesupper, strings.ToUpper(nome))
	}
	return nomesupper
}

func NomesToLowerGET(nomes []string) []string {
	nomeslower := []string{}
	for _, nome := range nomes {
		nomeslower = append(nomeslower, strings.ToLower(nome))
	}
	return nomeslower
}
