package main

import "fmt"

func main() {
	var pessoas = make([]map[string]string, 0)

	pessoas = append(pessoas, map[string]string{
		"nome":      "Joao",
		"sobrenome": "Silva"})
	pessoas = append(pessoas, pessoaADD("Zé", "Silva"))
	pessoas = append(pessoas, pessoaADD("Maria", "Silva"))
	for _, pessoa := range pessoas {
		fmt.Printf("%v, %v\n", pessoa["nome"], pessoa["sobrenome"])
	}
}

func pessoaADD(nome string, sobrenome string) map[string]string {
	var pessoa = make(map[string]string)
	pessoa["nome"] = nome
	pessoa["sobrenome"] = sobrenome
	return pessoa
}
