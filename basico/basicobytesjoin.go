package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := [][]byte
	{
	 []byte("um"), 
	 []byte("dois"), 
	 []byte("tres")
	}

	bytes.Join(s, []byte(", "))
	fmt.Printf("%s", s)

}
