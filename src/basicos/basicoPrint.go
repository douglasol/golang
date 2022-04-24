package basicos

import "fmt"

func Tprint() {
	nome := "Zé"

	fmt.Println(nome)
	fmt.Printf("Hello word %v\n", nome)
	print(nome)

	saida := fmt.Sprint(nome)
	print(saida)
}
