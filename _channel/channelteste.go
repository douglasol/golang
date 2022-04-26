package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go task(done)
	<-done
}

func task(done chan bool) {
	fmt.Println("Start Task")
	time.Sleep(time.Second * 3)
	fmt.Println("End Task")
	done <- true
}
