package main

import (
	"exemplo4/nomes"
	"fmt"
)

func main() {

	listanomes := nomes.NomesGET()
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}

	listanomes = nomes.NomesToUpperGET(listanomes)
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}
}
