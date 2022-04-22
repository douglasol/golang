/* funções em separado */
package main

import "fmt"

func main() {

	alomamae()

	nomes := nomesGET()
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}

	nomes = nomesToUpperGET(nomes)
	for _, nome := range nomes {
		fmt.Printf("%v\n", nome)
	}
}
