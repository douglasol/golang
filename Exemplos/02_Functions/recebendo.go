package main

import "fmt"

func main() {
	nomes := nomesGET()
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
