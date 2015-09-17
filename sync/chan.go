package main

import (
	"fmt"
	"sync"
)


const (
	N_INCREAMENTS = 1000000
)

func main() {
	counter := 0
	donechan := make(chan bool)
	var 

	go func(done chan<- bool) {
		for i := 0; i < N_INCREAMENTS; i++ {
			counter++
		}

		done <- true
	}(donechan)

	for i := 0; i < N_INCREAMENTS; i++ {
		counter++
	}

	_ = <-donechan

	fmt.Println("Counter: ", counter)
}
