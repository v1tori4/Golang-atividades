//Concorrência – 2. Goroutines & WaitGroups

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

//main
func main() {

	fmt.Println(runtime.NumCPU())       //mostrar o numero de processadores; executa um pouquinho de cada, em sequencia
	fmt.Println(runtime.NumGoroutine()) //começa o programa com uma goroutine (função main)

	wg.Add(1)

	go func1() //roda essas 2 funções
	go func2()

	fmt.Println(runtime.NumGoroutine()) //antes das funções serem executadas, executa isso

	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done()
}
