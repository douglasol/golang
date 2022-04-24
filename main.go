package main

import (
	"fmt"
	"golang/src/basicos"
)

func main() {

	var c int
	var t int

	fmt.Println("1. basicos")
	fmt.Println("2. functions")
	fmt.Println("3. package")
	fmt.Print("Qual o capitulo?")
	fmt.Scan(&c)

	fmt.Print("Qual o programa?")
	fmt.Scan(&t)

	switch c {
	case 1:
		// 1. basicos
		switch t {
		case 1:
			basicos.Tarray()
		case 2:
			basicos.Tprint()
		case 3:
			basicos.Tscan()
		case 4:
			basicos.Tregex()
		case 5:
			basicos.TconcatString()
		case 6:
			basicos.Tslice()
		case 7:
			basicos.Tfor()
		}

	}

}
