//Condicionais: a declaração if
//operador "!": negação

package main

import "fmt"

//main
func aa() {
	if x := 10; !(x < 100) {
		fmt.Println("X menor que 100")
	} else {
		fmt.Println("X maior que 100")
	}
}
