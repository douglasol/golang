package main

import (
	"encoding/ascii85"
	"fmt"
	"io/ioutil"
)

func main() {

	// byte to string
	var palavra string
	var bytes1 = []byte{'o', 'l', 'á', '!'}
	palavra = string(bytes1)
	fmt.Println(palavra)
	for i := 0; i < len(bytes1); i++ {
		fmt.Printf("%c", bytes1[i])
	}

	// ascii85 - pdf files
	fmt.Println("---")
	stringenc := make([]byte, 25, 25)
	stringdec := make([]byte, 25, 25)
	ascii85.Encode(stringenc, []byte("Olá, mamãe!"))
	ascii85.Decode(stringdec, stringenc, false)
	fmt.Println(string(stringdec))

	// string to byte
	var palavra2 = string("ola mamae!")
	var bytes2 = []byte("ola mamae!")
	var bytes3 = []byte(palavra2)
	fmt.Println(bytes2)
	fmt.Println(bytes3)

	// file to string
	bytes4, _ := ioutil.ReadFile("./basico/arquivo.txt")
	arquivo := string(bytes4)
	fmt.Println("arquivo:", arquivo)

}
