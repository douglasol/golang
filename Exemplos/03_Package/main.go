/* packages */

package main

import (
	"fmt"
	"package03/nomes"
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
