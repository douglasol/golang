/* JSON */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type pessoaDATA struct {
		Cpf   string `json:"id"`
		Nome  string `json:"-"`
		Email string `json:"email,omitempty"`
	}

	pessoa := pessoaDATA{
		Cpf:   "111",
		Nome:  "John Doe",
		Email: "",
	}

	/* converte em JSON */
	b, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(b)
	}
}
