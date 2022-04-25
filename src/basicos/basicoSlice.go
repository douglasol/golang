package basicos

import "fmt"

func Tslice() {
	sequencia  := []string{}
	sequencia = append(sequencia, "primeiro")
	sequencia = append(sequencia, "segundo")
	sequencia = append(sequencia, "terceiro")
	fmt.Printf("Valores = %v\n", sequencia)

}
''