// Ponteiros – 1. O que são ponteiros?
// - Todos os valores ficam armazenados na memória.
// - Toda localização na memória possui um endereço.
// - Um pointeiro se refere a esse endereço.

package main

import "fmt"

//main
func aa() {

	x := 0

	y := &x //endereco x

	fmt.Print(x, y)

	*y++ //pegar o valor que tem no endereco x que vai dar um

	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Print(x, y)
}
