package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var contador int

const quantidadeDeGourotines = 5000

func main() {

	criarGoroutines(quantidadeDeGourotines)
	wg.Wait()
	fmt.Println("Total de gourotines:\t", quantidadeDeGourotines,
		"\nTotal do contador:\t", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}

}

// Total de gourotines:     5000
// Total do contador:       5000
