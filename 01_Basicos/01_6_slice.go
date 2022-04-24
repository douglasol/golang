package main

import "fmt"

func tSlice() {
	//var booking []string
	booking := []string{}
	booking = append(booking, "primeiro")
	booking = append(booking, "segundo")
	booking = append(booking, "terceiro")
	fmt.Printf("Valores = %v\n", booking)

}
