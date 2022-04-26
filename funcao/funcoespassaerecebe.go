package main

import (
	"fmt"
	"golang/src/functions/funcs"
)

func main() {
	nomes := funcs.NomesGET()
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}

	nomes = funcs.NomesToUpperGET(nomes)
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
}
