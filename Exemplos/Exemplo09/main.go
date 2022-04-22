/* JSON */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type pessoaDATA struct {
		Cpf   string
		Nome  string
		Email string
	}

	pessoa := pessoaDATA{
		Cpf:   "111",
		Nome:  "John Doe",
		Email: "gardener@email.com",
	}

	/* converte em JSON */
	b, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(b)
	}
}
