package main

import "fmt"

func tScan() {
	var nome string
	fmt.Print("Forneça seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Ola %v", nome)
}
