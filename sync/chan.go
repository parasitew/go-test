package main

// This is a git test

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
	var m sync.Mutex

	go func(done chan<- bool) {
		for i := 0; i < N_INCREAMENTS; i++ {
			m.Lock()
			counter++
			m.Unlock()
		}

		done <- true
	}(donechan)

	for i := 0; i < N_INCREAMENTS; i++ {
		m.Lock()
		counter++
		m.Unlock()
	}

	_ = <-donechan

	fmt.Println("Counter: ", counter)
}
