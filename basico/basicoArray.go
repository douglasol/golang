package main

import "fmt"

func main() {

	var booking [3]string
	booking[0] = "primeiro"
	booking[1] = "segundo"
	booking[2] = "terceiro"

	fmt.Printf("Tamanho = %v\n", len(booking))
	fmt.Printf("Tipo = %T\n", booking)

	for i, item := range booking {
		fmt.Printf("booking[%v]=%v\n", i, item[i])
	}
}
