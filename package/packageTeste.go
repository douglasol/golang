package package

import (
	"fmt"
	"golang/package/pck1"
)

func main() {

	listanomes := pcknome.NomesGET()
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}

	listanomes = pcknome.NomesToUpperGET(listanomes)
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}

	listanomes = pcknome.NomesToLowerGET(listanomes)
	for _, nome := range listanomes {
		fmt.Printf("%v\n", nome)
	}
}
