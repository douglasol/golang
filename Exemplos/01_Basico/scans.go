package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Forneça seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Ola %v", nome)
}
