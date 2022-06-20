package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	novasgoroutines(10)
	wg.Wait()
}

func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}

// Eu sou a goroutine número: 10
// Eu sou a goroutine número: 5
// Eu sou a goroutine número: 2
// Eu sou a goroutine número: 3
// Eu sou a goroutine número: 4
// Eu sou a goroutine número: 7
// Eu sou a goroutine número: 6
// Eu sou a goroutine número: 8
// Eu sou a goroutine número: 9
// Eu sou a goroutine número: 1
