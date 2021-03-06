package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan bool)
	go hardTask(done)

	wg.Add(1)
	go loader(done)
	wg.Wait()
}

// nao espera valor no channel, somente escreve
func hardTask(done chan<- bool) {
	fmt.Println("start task")
	time.Sleep(time.Second * 3)
	done <- true
}

func loader(done <-chan bool) {
	defer wg.Done()
	i := 0
	load := []rune(`-/-|-\`)
	for {
		/* auxilia a controlar o channel, timeout, ... */
		select {
		case <-done:
			fmt.Printf("\r")
			fmt.Println("Done")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(load[i]))
			i++
			if i == len(load) {
				i = 0
			}

		}

	}
}
