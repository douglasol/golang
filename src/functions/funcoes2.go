package main

import (
	"fmt"
	"golang/src/funcoes"
)

func main() {
	nomes := funcoes.NomesGET()
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}

	nomes = funcoes.NomesToUpperGET(nomes)
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
}
