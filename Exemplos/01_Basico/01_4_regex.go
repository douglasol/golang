package main

import (
	"fmt"
	"regexp"
)

func tRegex() {
	var inome string
	var isok bool

	fmt.Print("Nome (. Encerra): ")
	fmt.Scan(&inome)

	/* strings vazias */
	match1, _ := regexp.MatchString("^$", inome)
	fmt.Printf("vazia: %v\n", match1)

	/* tamanho da string */
	isok = len(inome) >= 2
	fmt.Printf("tamanho: %v\n", isok)

	// somente letras e espaço
	match2, _ := regexp.MatchString("/^[a-zA-Z]*$/", inome)
	fmt.Printf("somente letras e espaço: %v\n", match2)

	// somente numeros
	match3, _ := regexp.MatchString("^[0-9]", inome)
	fmt.Printf("somente numeros: %v\n", match3)
}
