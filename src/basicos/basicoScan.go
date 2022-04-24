package basicos

import "fmt"

func Tscan() {
	var nome string
	fmt.Print("Forneça seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Ola %v", nome)
}
