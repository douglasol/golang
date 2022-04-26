package main

//https://www.youtube.com/watch?v=4C8l7eETr24
import (
	"fmt"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)

	go sendMessage(msg)
	go receiveMessage(msg, done)

	<-done
}

func sendMessage(msg chan string) {
	msg <- "Ola, mensagem sendo enviada"
}

func receiveMessage(msg chan string, done chan bool) {
	fmt.Println(<-msg)
	done <- true
}
