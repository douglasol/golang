package main

import (
	"fmt"
	"modulo/pacote/pacote1"
)

func main() {

	listanomes := pacote1.NomesGET()
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}

	listanomes = pacote1.NomesToUpperGET(listanomes)
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}

	listanomes = pacote1.NomesToLowerGET(listanomes)
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}
}
