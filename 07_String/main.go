package main

/* string */
import "fmt"

func main() {
	texto := "Isso é uma string"
	fmt.Printf("%v %v %T", string(texto[2]), texto[2], texto[2])

	/* incluir:
	 	strings.TrimSpace()
		[]byte to string
	*/
}
