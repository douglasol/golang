package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./file/arquivo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 256)
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("%s", hex.Dump(buffer))
	}
}
