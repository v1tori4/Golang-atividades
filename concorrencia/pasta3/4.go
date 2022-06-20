//Concorrência – 5. Mutex
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	//var mu sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			//mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}

//com
// CPUs: 4
// Goroutines: 1
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 1
// Valor final: 15

//sem
// CPUs: 4
// Goroutines: 1
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 2
// Goroutines: 1
// Valor final: 15
