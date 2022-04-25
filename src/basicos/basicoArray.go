package basicos

import "fmt"

func Tarray() {

	var sequencia [3]string
	sequencia[0] = "primeiro"
	sequencia[1] = "segundo"
	sequencia[2] = "terceiro"

	fmt.Printf("Tamanho = %v\n", len(sequencia))
	fmt.Printf("Tipo = %T\n", sequencia)

	for i, item := range sequencia {
		fmt.Printf("sequencia [%v]=%v\n", i, item[i])
	}
}
