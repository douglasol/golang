package main

import (
	"fmt"
	"log"

	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
)

type Pessoa struct {
	Cpf   string
	Nome  string
	Email string
}

func main() {
	println("TEM PROBLEMA NO PREPARE, NAO ESTA ENTENDENDO O ? ")
	isOk := mydb.MyDbConnect()
	if isOk {
		pessoas := []Pessoa{
			{"123123", "Jardel Silva", "jardel@j.com"},
			{"234234", "Jorel Silva", "jorel@j.com"},
			{"345345", "Janio Silva", "janio@j.com"},
			{"456456", "Junio Antonio", "junio@j.com"},
		}
		total, err := pessoaINSERTS(pessoas)
		if err != nil {
			log.Fatal("Error reading Pessoa: ", err.Error())
		}
		fmt.Printf("inseridos %d registro(s) com sucesso.\n", total)
	} else {
		fmt.Printf("Não! %v", isOk)
	}
}

func pessoaINSERTS(pessoas []Pessoa) (int, error) {
	var total int

	err, _ := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}

	stmt, err := mydb.Db.Prepare("INSERT INTO pessoa (PessoaCpf, PessoaNome, PessoaEmail) VALUES(?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	for _, pessoa := range pessoas {
		if _, err := stmt.Exec(pessoa.Cpf, pessoa.Nome, pessoa.Email); err != nil {
			log.Fatal(err)
		}
		total++
	}

	return total, nil
}
