package main

import "fmt"

func main() {
	var pessoa = map[string]int{
		"Zé":    10,
		"Maria": 20,
	}
	fmt.Printf("%v\n", pessoa)

	var menu map[string]float64
	menu = map[string]float64{
		"ovo":      1.75,
		"bacon":    3.22,
		"salsicha": 1.89,
	}
	fmt.Printf("%v\n", menu)
}
