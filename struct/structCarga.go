package main

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

func main() {

	// inicialização simples
	pessoa := Pessoa{"111.111.111-11", "Jao Gard", "gardener@email.com"}

	// inicialização de coleção
	pessoas := []Pessoa{
		{"111.111.111-11", "Jao Gard", "gardener@email.com"},
		{"222.222.222-22", "Rogerio Driv", "driver@email.com"},
		{"333.333.333-33", "Pragma Dor", "programmer@email.com"},
	}

	// inicialização
	pessoas = []Pessoa{
		{Cpf: "444.444.444-44", Nome: "Zé", Email: "zz@zz.com"},
		{Cpf: "555.555.555-55", Nome: "Maria", Email: "mm@mm.com"},
	}

	// com append
	pessoas = append(pessoas, Pessoa{"666.666.666-66", "Um", "ze@ze.com"})

	// com variáveis
	var nome, cpf, email string
	nome = "Carlos"
	cpf = "777.777.777-77"
	email = "cc@cc.com"
	pessoa := Pessoa{cpf, nome, email}
	pessoas = append(pessoas, pessoa)

	// com uma função
	pessoas = append(pessoas, pessoaADD("888.888.888-88", "Mané", "ze@ze.com"))
}

func pessoaADD(cpf string, nome string, email string) Pessoa {
	var pessoa = Pessoa{
		Cpf:   cpf,
		Nome:  nome,
		Email: email,
	}
	return pessoa
}
