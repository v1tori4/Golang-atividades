package main

import (
	"fmt"
	"runtime"
	"sync"
)

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
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}

}

// Total de gourotines:     5000
// Total do contador:       480
