package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Pessoas struct {
	Colecao []struct {
		Cpf   string `json:"cpf"`
		Nome  string `json:"nome"`
		Email string `json:"email"`
	} `json:"colecao"`
}

func main() {
	file, _ := ioutil.ReadFile("./json/colecao.json")
	pessoas := Pessoas{}
	_ = json.Unmarshal([]byte(file), &pessoas)
	for i := 0; i < len(pessoas.Colecao); i++ {
		fmt.Printf("%v\n", pessoas.Colecao[i].Cpf)
	}
}
