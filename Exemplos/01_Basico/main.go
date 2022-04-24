package main

import "fmt"

func main() {

	var i int
	fmt.Print("Qual o programa?")
	fmt.Scan(&i)

	switch i {
	case 1:
		tArray()
	case 2:
		tPrint()
	case 3:
		tScan()
	case 4:
		tRegex()
	case 5:
		tConcatString()
	case 6:
		tSlice()
	case 7:
		tFor()
	}
}
