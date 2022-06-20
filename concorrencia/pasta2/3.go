//Concorrência – 3. Discussão: Condição de corrida

// *Função 1       var     Função 2*
// Lendo: 0   →   0
// Yield          0   →   Lendo: 0
// var++: 1               Yield
// Grava: 1   →   1       var++: 1
// 			   1   ←   Grava: 1
// Lendo: 1   ←   1
// Yield          1   →   Lendo: 1
// var++: 2               Yield
// Grava: 2   →   2       var++: 2
// 			   2   ←   Grava: 2

//Na prática: Condição de corrida

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//main
func main() {

	fmt.Println("CPUs:", runtime.NumCPU()) //ver quantos processadores tem
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	contador := 0
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched() //yield
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
