package main

import (
	"fmt"
	"strings"
)

func main() {
	nomes, nomesupper, nomeslower := nomesListGET()
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
	for _, nome := range nomesupper {
		fmt.Printf("%v\n", nome)
	}
	for _, nome := range nomeslower {
		fmt.Printf("%v\n", nome)
	}
}

func nomesListGET() ([]string, []string, []string) {
	nomes := []string{}
	nomesupper := []string{}
	nomeslower := []string{}
	// carga
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	// conversao maiusculo
	for _, nome := range nomes {
		nomesupper = append(nomesupper, strings.ToUpper(nome))
	}
	// conversao minusculo
	for _, nome := range nomes {
		nomeslower = append(nomeslower, strings.ToLower(nome))
	}
	return nomes, nomesupper, nomeslower
}
