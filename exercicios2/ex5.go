//mostrar resto da divisao por 4 de nuemro de 10 a 100
//operador modulo

package main

import "fmt"

//main
func resto() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}

}
