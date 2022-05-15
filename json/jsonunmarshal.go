package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type pessoa struct {
	Cpf   string `json: "cpf"`
	Nome  string `json: "nome"`
	Email string `json: "email"`
}

type pessoas struct {
	Pessoas []pessoa
}

func main() {
	file, _ := ioutil.ReadFile("./json/arquivo.json")
	data := pessoa{}
	_ = json.Unmarshal([]byte(file), &data)
	fmt.Printf("%v %v %v ", data.Cpf, data.Email, data.Nome)
}
