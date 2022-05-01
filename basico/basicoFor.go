package main

import (
	"fmt"
	"strings"
)

func main() {
	nomecompleto := []string{}
	primeironome := []string{}

	for {
		var inome string
		var isobrenome string
		fmt.Println("Forneça seu nome (. encerrra):")
		fmt.Scan(&inome)

		continuar := inome == "."
		if continuar {
			break
		} else {
			fmt.Println("Forneça seu sobrenome:")
			fmt.Scan(&isobrenome)
			nomecompleto = append(nomecompleto, inome+" "+isobrenome)
		}
	}
	for _, nome := range nomecompleto {
		var nomes = strings.Fields(nome)
		primeironome = append(primeironome, nomes[0])
	}

	for i, nome := range primeironome {
		fmt.Printf("Nome[%v]=%v\n", i, nome)
	}
}
